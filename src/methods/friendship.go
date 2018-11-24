package methods

import (
	"../constants"
	"../entities"
	"../structures"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"log"
)

var friendship = createMethod(runFriendship, friendShipMethodSettings)

func runFriendship(request structures.Request) {
	if request.Namespace == constants.NAMESPACE_INTERNAL {
		return
	}
	HandshakeMsg.Id = generateId()
	sendToInternal(HandshakeMsg)
}

func sendToInternal(request structures.Request) {
	_requestByte, marshalErr := json.Marshal(request)
	FailOnError(marshalErr, "Failed on marshal request message.")

	err := entities.Rabbit.Channels[constants.NAMESPACE_INTERNAL].Publish(
		"",     // exchange
		constants.NAMESPACE_INTERNAL, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:        []byte(_requestByte),
		})
	FailOnError(err, "Failed to publish a message to internal service.")
	log.Printf("%s Sent message to [* %s *]: Message %s", constants.HEADER_RMQ_MESSAGE, constants.NAMESPACE_INTERNAL, _requestByte)
}

var friendShipMethodSettings = structures.MethodSettings {
	IsInternal: true,
	Auth: false,
	Cache: 0,
	Middlewares: structures.Middlewares {
		Before: []string{},
		After: []string{},
	},
}

var handshakeParams = map[string] interface {} {
	"namespace": constants.CONFIG.Namespace,
	"methods": map[string] interface {} {
		"friendship": friendShipMethodSettings,
		"infrastructure": infrastructureMethodSettings,
	},
}

var HandshakeMsg = structures.Request {
	Namespace: constants.NAMESPACE_INTERNAL,
	Method: "handshake",
	Domain: nil,
	Locale: nil,
	Params: handshakeParams,
	Source: constants.CONFIG.Namespace,
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func generateId() uuid.UUID {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong with generate id: %s", err)
		return uuid.UUID{}
	}
	return id
}