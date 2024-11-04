package amqp

type subscribeOptions struct {
	topic    string
	dataType string
}

var defaultSubscribeOptions = subscribeOptions{}

type SubscribeOption interface {
	apply(*subscribeOptions)
}

type funcSubscribeOptions struct {
	f func(*subscribeOptions)
}

func (f *funcSubscribeOptions) apply(o *subscribeOptions) {
	f.f(o)
}

func newFuncSubscribeOptions(f func(*subscribeOptions)) SubscribeOption {
	return &funcSubscribeOptions{f: f}
}

// WithType set the subscriber to listen on topic (& typ) only
// typ => data type if supplied, only the first one will take effect
// P.S. the topic below must refer to client option's queue(s)
// If topic if empty string, typ is not supplied => will subscribe to all topic
// If topic & typ is empty string => will subscribe to all topic & data type
// If topic is "topic-name", typ is not supplied => will subscribe to "topic-name"'s all data type
// If topic is "", typ is "Type1" => will subscribe to all topic with "Type1" data type
// If topic is "topic-name", typ is "" => will subscribe to topic "topic-name"'s all data type
func WithSubscribeDataType(topic string, typ string) SubscribeOption {
	return newFuncSubscribeOptions(func(o *subscribeOptions) {
		o.topic = topic
		o.dataType = typ
	})
}