package src

import (
	"./constants"
	"./templates"
	"encoding/json"
	"fmt"
	"log"
)

func getMessageHeader(t string) string {
	switch t {
	case "rmq":
		return constants.HEADER_RMQ_MESSAGE
	case "http":
		return constants.HEADER_HTTP_MESSAGE
	case "ws":
		return constants.HEADER_WS_MESSAGE
	default:
		return constants.HEADER_UNKNOWN
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func logRequest(request templates.Request, t string) {
	fmt.Printf("%s Get request %+v\n", getMessageHeader(t), request)
}

func unMarshalMessage(data []byte, variable *map[string] interface{}) {
	if err := json.Unmarshal(data, &*variable); err != nil {
		FailOnError(err, "Error on parse message")
	}
}

func setSource(request *templates.Request, value string) {
	request.Source = value
}
