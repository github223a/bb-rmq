package src

import (
	"./constants"
	"./structures"
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
	case "redis":
		return constants.HEADER_REDIS_MESSAGE
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

func FailOnError(err error, msg string, source string) {
	header := getMessageHeader(source)
	if err != nil {
		log.Fatalf("%s %s: %s", header, msg, err)
	}
}

func logRequest(request structures.Request, t string) {
	fmt.Printf("%s Get request %+v\n", getMessageHeader(t), request)
}

func setSource(request *structures.Request, value string) {
	request.Source = value
}
