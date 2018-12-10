package src

import (
	rmq "bb_rmq"
	"fmt"
	"log"
)

func getMessageHeader(t string) string {
	switch t {

	case "http":
		return HEADER_HTTP_MESSAGE
	case "ws":
		return HEADER_WS_MESSAGE
	case "redis":
		return HEADER_REDIS_MESSAGE
	default:
		return HEADER_UNKNOWN
	}
}

func getNamespaceSettings(request rmq.Request) rmq.NamespaceSettings {
	return InfrastructureData.Infrastructure[request.Namespace]
}

func getMethodSettings(request rmq.Request) rmq.MethodSettings {
	return getNamespaceSettings(request).Methods[request.Method]
}

func FailOnError(err error, msg string, source string) {
	header := getMessageHeader(source)
	if err != nil {
		log.Fatalf("%s %s: %s", header, msg, err)
	}
}

func logRequest(request rmq.Request, t string) {
	fmt.Printf("%s Get request %+v\n", getMessageHeader(t), request)
}

func setSource(request *rmq.Request, value string) {
	request.Source = value
}
