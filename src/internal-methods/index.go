package internal_methods

import "github.com/streadway/amqp"

var List = map[string] func(*amqp.Channel) {
	"friendship": handshake,
	"infrastructure": infrastructure,
}
