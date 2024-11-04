package amqp

import "time"

const (
	defaultTimeout  = time.Minute
	reConnectDelay  = 2 * time.Second
	maxMessageChannelSize = 20
	defaultClientIDSize = 12

	// Temporary queue will remove after current connection disconnect
	// Important: If temporary queue is not empty, will be retained.
	Temporary lifeCycle = iota
	// Exclusive queue will only be used by the current connection. Any attempt to use this queue
	// from a different connection will result in a channel-level exception `RESOURCE_LOCKED`.
	// This queue lifeCycle will delete itself after connection lost (actively/passively disconnect)
	Exclusive
	// Durable queue will not delete after disconnect(client) or restart(server)
	Durable
)
var defaultOptions = options{
	queue:              make([]Queue, 0),
}