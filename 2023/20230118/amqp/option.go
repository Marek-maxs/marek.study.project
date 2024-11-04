package amqp

// Option represent the client options
type Option interface {
	apply(*options)
}

type options struct {
	block              bool
	broker             Broker
	topic              string
	queue              []Queue
	pubDefaultExchange Exchange
}

// Broker is the broker's setting
type Broker struct {
	user    string
	passwd  string
	host    string
	port    int
	sslmode bool // only for AMQPS
}

type lifeCycle int

// Queue represent a message queue
type Queue struct {
	// Name of the queue
	Name string
	// Name of the topic
	Topic string
	// Lifecycle represent the queue's life cycle
	Lifecycle lifeCycle
	// ExchangeName should be defiend before using it.
	// Exchange indicate the rule how message will be dispatched to queue.
	// This option is only used when the queue without binding.
	ExchangeBind Exchange
}