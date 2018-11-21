package internal_methods

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

type Empty struct {}

var handshakeParams = map[string] interface {} {
	"namespace": constants.CONFIG.Namespace,
	"methods": Empty{},
}

var handshakeRequest = templates.Request {
	Id: generateId(),
	Namespace: constants.NAMESPACE_INTERNAL,
	Method: "handshake",
	Domain: nil,
	Locale: nil,
	Params: handshakeParams,
	Source: constants.CONFIG.Namespace,
}

func handshake(request templates.Request) {
	var handshakeMsgByte, marshalErr = json.Marshal(handshakeRequest)
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
}
