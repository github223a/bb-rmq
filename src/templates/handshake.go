package templates

import (
	"fmt"
	"github.com/satori/go.uuid"
)

const NAMESPACE = "registration"
const HANDSHAKE = "handshake"
const INTERNAL = "internal"

func generateId() uuid.UUID{
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return uuid.UUID{}
	}
	return id
}

type Empty struct {}

type Params struct {
	Namespace string `json:"namespace"`
	Methods Empty `json:"methods"`
}

type Request struct {
	Id uuid.UUID `json:"id"`
	Namespace string `json:"namespace"`
	Method string `json:"method"`
	Domain string `json:"domain"`
	Locale string `json:"locale"`
	Params Params `json:"params"`
	Source string `json:"source"`
}




var req = Request{
	generateId(),
	INTERNAL,
	HANDSHAKE,
	"russia",
	"ru",
	Params{
		Namespace: NAMESPACE,
		Methods: Empty{},
	},
	NAMESPACE,
}

func Handshake() Request{
	return req
}
