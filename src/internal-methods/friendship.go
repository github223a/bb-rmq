package internal_methods

import (
	"../constants"
	"../templates"
	"fmt"
	"github.com/satori/go.uuid"
)

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
	"namespace": constants.NAMESPACE_INTERNAL,
	"method": Empty{},
}

var handshakeRequest = templates.Request {
	Id: generateId(),
	Namespace: constants.CONFIG.Namespace,
	Method: "handshake",
	Domain: "",
	Locale: "",
	Params: handshakeParams,
	Source: constants.CONFIG.Namespace,
}

func handshake() {
	fmt.Println(handshakeRequest)
}
