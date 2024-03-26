package amqp

import (
	"github.com/en-trak/mqclient/v3"
	"io"
)

type Subscriber interface {
	io.Closer
	IsConnected() bool
	Subscribe(opt ...SubscribeOption) (<-chan Message, error)
}

type client struct {
	c *connection
	opts options
}

// NewSubscriber will create a AMQP 0.9.1 task queue subscriber
func NewSubscriber(opt ...Option) (Subscriber, error) {
	opts := defaultOptions
	for _, o := range opt {
		o.apply(&opts)
	}

	conn, err := newConnection(opts.broker, opts.block, opts.queue)
	if err != nil {
		return nil, err
	}

	return &client{c: conn, opts: opts}, nil
}

func (c *client) Subscribe(opt ...SubscribeOption) (<-chan Message, error) {
	if !c.c.isReady() {
		return nil, mqclient.ErrConnectionNotReady
	}

	opts := defaultSubscribeOptions

	for _, o := range opt {
		o.apply(&opts)
	}

	return c.c.addSubscriber(opts.topic, opts.dataType)
}

func (c *client) IsConnected() bool {
	return !c.c.isClose()
}

func (c *client) Close() error {
	return c.c.close()
}