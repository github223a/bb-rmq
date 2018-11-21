package entities

import "github.com/streadway/amqp"

var Rabbit *Rmq

type Rmq struct {
	Connection *amqp.Connection
	Channels map[string] *amqp.Channel
}

func (rmq *Rmq) SetConnection (connection *amqp.Connection) {
	rmq.Connection = connection
}

func (rmq *Rmq) AddChannel (name string, channel *amqp.Channel) {
	rmq.Channels[name] = channel
}

func NewRabbitEntity(connection *amqp.Connection) *Rmq {
	var channels = make(map[string] *amqp.Channel)
	return &Rmq{
		Connection: connection,
		Channels: channels,
	}
}
