package amqp

type Exchange string

func (e Exchange) String() string {
	return string(e)
}

const (
	// Will use direct as default when exchange is empty, the same as ExchangeDefault
	// Deprecated: Try not to use it and use ExchangeDirect instead.
	ExchangeEmpty Exchange = ""
	// ExchangeDefault is default exchange of direct type, the same as ExchangeEmpty
	// try not to use it.
	ExchangeDefault Exchange = "amq.default"
	// ExchangeDirect delivers messages to queues based on the message routing key.
	// ExchangeDirect is ideal for the unicast routing of messages
	// (although they can be used for multicast routing as well).
	ExchangeDirect Exchange = "amq.direct"
	// ExchangeFanout routes messages to all of the queues that
	// are bound to it and the routing key is ignored.
	// If N queues are bound to a fanout exchange,
	// when a new message is published to that
	// exchange a copy of the message is delivered to all N queues.
	// Fanout exchanges are ideal for the broadcast routing of messages.
	ExchangeFanout Exchange = "amq.fanout"
	// ExchangeHeaders is designed for routing on multiple attributes that
	// are more easily expressed as message headers than a routing key.
	// Headers exchanges ignore the routing key attribute. Instead,
	// the attributes used for routing are taken from the headers attribute.
	ExchangeHeaders Exchange = "amq.headers"
	// ExchangeMatch another type of ExchangeHeaders
	ExchangeMatch Exchange = "amq.match"
	// ExchangeTopic route messages to one or many queues based on matching
	// between a message routing key and the pattern
	// that was used to bind a queue to an exchange.
	ExchangeTopic Exchange = "amq.topic"
)
