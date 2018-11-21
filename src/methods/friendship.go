package methods

import (
	"../constants"
	"../entities"
	"../templates"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func generateId() uuid.UUID {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return uuid.UUID{}
	}
	return id
}

type HandshakeMethods struct {}

var handshakeParams = map[string] interface {} {
	"namespace": constants.CONFIG.Namespace,
	"methods": map[string] interface{} {
		"friendship": friendShipSettings,
		"infrastructure": infrastructureSettings,
	},
}

var HandshakeRequest = templates.Request {
	Id: generateId(),
	Namespace: constants.NAMESPACE_INTERNAL,
	Method: "handshake",
	Domain: nil,
	Locale: nil,
	Params: handshakeParams,
	Source: constants.CONFIG.Namespace,
}

func runFriendship(request templates.Request) {
	handshakeMsgByte, marshalErr := json.Marshal(HandshakeRequest)
	FailOnError(marshalErr, "Failed on marshal handshake message.")

	err := entities.Rabbit.Channels[request.Namespace].Publish(
		"",     // exchange
		constants.NAMESPACE_INTERNAL, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:        []byte(handshakeMsgByte),
		})
	FailOnError(err, "Failed to publish a message.")
	log.Printf("%s Sent message to [*%s*]. Message %s", constants.HEADER_RMQ_MESSAGE, constants.NAMESPACE_INTERNAL, handshakeMsgByte)
}

var friendShipSettings = entities.MethodSettings{
	IsInternal: true,
	Auth: false,
	Cache: nil,
	Middlewares: entities.Middlewares {
		Before: []string{},
		After: []string{},
	},
}

var Friendship = entities.NewMethodEntity(runFriendship, friendShipSettings)
