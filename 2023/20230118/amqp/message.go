package amqp

import (
	"fmt"
	gamqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

type Message interface {
	// ID return the message id
	ID() string
	// Topic return the current message's topic
	Topic() string
	// ContentType return the message content-type
	ContentType() string
	// Payload return the content of the message
	Payload() []byte
	// Ack response ack to broker
	Ack() error
	// Nack response nack to broker
	Nack(requeue bool) error
}

type message struct {
	r gamqp.Delivery
}

func newMessage(d gamqp.Delivery) Message {
	return &message{r: d}
}

func (m *message) ID() string {
	return strconv.FormatUint(m.r.DeliveryTag, 10) // nolint
}

func (m *message) Topic() string {
	return fmt.Sprintf("%s; %s", m.r.RoutingKey, m.r.Type)
}

func (m *message) ContentType() string {
	return m.r.ContentType
}

func (m *message) Payload() []byte {
	return m.r.Body
}

func (m *message) Ack() error {
	return m.r.Ack(false)
}

func (m *message) Nack(requeue bool) error {
	return m.r.Nack(false, requeue)
}