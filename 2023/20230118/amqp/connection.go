package amqp

import (
	gamqp "github.com/rabbitmq/amqp091-go"
	gonanoid "github.com/matoous/go-nanoid/v2"
	mqclient "github.com/en-trak/mqclient/v3"
	"sync"
	"github.com/rs/zerolog/log"
	"time"
)

type connection struct {
	c	*gamqp.Connection
	url string

	queues []Queue
	delivers chan gamqp.Delivery

	// a sub-router of data type in client side
	subscribers map[chan Message]struct{
		RoutingKey	string
		DataType 	string
	}

	publishCh	*gamqp.Channel
	publishConfirm	chan gamqp.Confirmation
	publishReturn chan gamqp.Return

	mu 	sync.Mutex
	connMu  sync.Mutex
	once 	sync.Once
	done	chan struct{}
	notifyConnClose	chan *gamqp.Error
	ready 	bool
	closeChan	chan bool
}

func newConnection(broker Broker, block bool, queues []Queue) (*connection, error) {
	uri := &gamqp.URI{
		Scheme:     "amqp",
		Host:       broker.host,
		Port:       broker.port,
		Username:   broker.user,
		Password:   broker.passwd,
		Vhost:      "/",
	}

	if broker.sslmode {
		uri.Scheme += "s"
	}

	conn := &connection{
		url:uri.String(),
		done: make(chan struct{}),
		queues: queues,
	}

	log.Debug().Str("portocol", "AMQP").Str("host", conn.url).Msg("mqclient: add broker")

	go conn.handleReconnect()

	// might sub
	if len(queues) > 0 {
		go conn.listen()
	}

	if block {
		timer := time.NewTimer(defaultTimeout)
		
		for {
			if conn.isReady() {
				return conn, nil
			}
			
			select {
			case <-time.After(reConnectDelay):
			case <-timer.C:
				conn.close()
				return nil, mqclient.ErrConnectionTimeout
			}
		}
	}

	return conn, nil
}

func (c *connection) handleReconnect() {
	for {
		c.connMu.Lock()
		c.ready = false
		c.connMu.Unlock()

		err := c.connect()
		if err == nil {
			return
		}

		log.Error().Err(err).
			Str("protocol", "AMQP").
			Msg("mqclient: connect failed. retry in 2 seconds")

		select {
		case <-c.done:
			return
		case <-time.After(reConnectDelay):
		}
	}
}

func (c *connection) connect() error {
	conn, err := gamqp.Dial(c.url)
	if err != nil {
		return err
	}

	c.notifyConnClose = make(chan *gamqp.Error)
	conn.NotifyClose(c.notifyConnClose)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// when connection is about subscriber
	for _, q := range c.queues {
		if q.Name == "" {
			continue // do not use empty queue name
		}

		durable, autoDelete, exclusive := explanLifeCycle(q)

		// if queue name doesn't exits, will create a new one.
		// if exists dlready,will compare the configs and return error when configs are different.
		_, err := ch.QueueDeclare(q.Name, durable, autoDelete, exclusive, false, nil)
		if err != nil {
			return err
		}

		log.Debug().Str("protocol", "AMQP").
			Str("queue", q.Name).Bool("auto_delete", autoDelete).Bool("exclusive", exclusive).
			Msg("mqclient: queue declared")

		// Try to bind queue with exchange.
		// If the queue has been bound with one exchange befor, will ignore this binding.
		err = ch.QueueBind(q.Name, q.Topic, q.ExchangeBind.String(), false, nil)
		if err != nil {
			return err
		}

		log.Debug().Str("protocol", "AMQP").
			Str("queue", q.Name).Str("exchange", q.ExchangeBind.String()).
			Msgf("mqclient: queue bind exchange with topic %s", q.Topic)

		if err != nil {
			conn.Close()
			return err
		}
	}

	c.mu.Lock()
	c.c = conn
	c.publishCh = nil
	c.mu.Unlock()

	c.connMu.Lock()
	c.ready = true
	c.connMu.Unlock()

	log.Debug().Str("protocol", "AMQP").Msg("mqclient: connected")

	select {
	case err := <-c.notifyConnClose:
		return err
	case <-c.done:
	}

	return nil
}

func (c *connection) listen() {
	c.mu.Lock()

	if c.delivers == nil {
		c.delivers = make(chan gamqp.Delivery, maxMessageChannelSize)
	}

	if c.subscribers == nil {
		c.subscribers = make(map[chan Message]struct{
			RoutingKey 	string
			DataType 	string
		})
	}

	c.mu.Unlock()

	for {
		if len(c.subscribers) > 0 {
			break
		}

		time.Sleep(time.Second)
	}

	for d := range c.delivers {
		found := false

		c.mu.Lock()
		for ch, s := range c.subscribers {
			if checkTopic(d.RoutingKey, s.RoutingKey) && (s.DataType == d.Type || s.DataType == "") {
				ch <- newMessage(d)

				found = true
			}
		}
		c.mu.Unlock()

		if !found {
			log.Warn().
				Str("protocol", "AMQP").
				Str("topic", d.RoutingKey).
				Str("dataType", d.Type).
				Str("contentType", d.ContentType).
				Msg("mqclient: topic missing subscriber")
		}
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	for ch := range c.subscribers {
		close(ch)
	}

	c.subscribers = nil
}

func (c *connection) subscribe(conn *gamqp.Connection, queue Queue) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	var args gamqp.Table
	var noWait, autoAck, cExclusive, noLocal bool

	delivers, err := ch.Consume(queue.Name, gonanoid.Must(defaultClientIDSize),
		autoAck, cExclusive, noLocal, noWait, args)
	if err != nil {
		return err
	}
	go func() {
		defer ch.Close()
		for {
			select {
			case <-c.closeChan:
				return
			case d := <- delivers:
				c.delivers <- d
			}
		}
	}()

	return nil
}

func (c *connection) addSubscriber(routingKey, dataType string) (<- chan Message, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, queue := range c.queues {
		topic := queue.Topic
		if topic == routingKey || routingKey == "" {
			ch := make(chan Message, maxMessageChannelSize)
			c.subscribers[ch] = struct {
				RoutingKey string
				DataType   string
			}{
				RoutingKey: routingKey,
				DataType: dataType,
			}

			return ch, nil
		}
	}
	return nil, &mqclient.ErrQueueNotDeclared{}
}

func (c *connection) isReady() bool {
	c.connMu.Lock()
	defer c.connMu.Unlock()

	return c.ready
}

func (c *connection) isClose() bool {
	return c.c.IsClosed()
}

func (c *connection) close() error  {
	var err error
	c.once.Do(func() {
		close(c.done)
		if c.delivers != nil {
			close(c.delivers)
		}

		if c.c != nil {
			err = c.c.Close()
		}

		log.Debug().Str("protocol", "AMQP").Msg("mqclient: disconnected")
	})

	return err
}