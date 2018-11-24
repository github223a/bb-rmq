package src

import (
	"./constants"
	"./structures"
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

func getNamespaceSettings(request structures.Request) structures.NamespaceSettings {
	return constants.InfrastructureData.Infrastructure[request.Namespace]
}

func getMethodSettings(request structures.Request) structures.MethodSettings {
	return getNamespaceSettings(request).Methods[request.Method]
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func logRequest(request structures.Request, t string) {
	fmt.Printf("%s Get request %+v\n", getMessageHeader(t), request)
}

func UnmarshalByteToMap(data []byte, variable *map[string] interface{}) {
	if err := json.Unmarshal(data, &*variable); err != nil {
		FailOnError(err, "Error on unmarshal byte message")
	}
}

func setSource(request *structures.Request, value string) {
	request.Source = value
}
