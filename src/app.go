package src

import (
	"./constants"
	"./entities"
	"./methods"
	"./structures"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func processingExternalMethod(request structures.Request) {
	fmt.Printf("%+v\n", request)
	methodSettings := constants.InfrastructureData.Infrastructure[request.Namespace].Methods[request.Method]
	cacheTimer := methodSettings.Cache

	if constants.CONFIG.UseCache == true && cacheTimer > 0 {
		//request.cacheKey = getCacheKey(request)
		sendCachedResponse(request)
		return
	}

	_request, marshalErr := json.Marshal(request)
	FailOnError(marshalErr, "Failed on marshal request message.")

	applyBeforeMiddlewares(request)
	err := entities.Rabbit.Channels[request.Namespace].Publish(
		"",     // exchange
		constants.NAMESPACE_INTERNAL, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:        []byte(_request),
		})
	FailOnError(err, "Failed to publish a message.")
	log.Printf("%s Sent message to [* %s *]: Message %s", constants.HEADER_RMQ_MESSAGE, constants.NAMESPACE_INTERNAL, _request)
}

func processingInternalMethod(message structures.Request) {
	validateRequest(message)
	checkNamespace(message)
	checkInternalMethod(message)
	checkToken(message)

	methods.List[message.Method].Run(message)
}

func validateRequest(request structures.Request) {
}

func checkNamespace(request structures.Request) {
	namespace := request.Namespace
	_, ok := constants.InfrastructureData.Infrastructure[namespace]

	if namespace != constants.NAMESPACE_INTERNAL && namespace != constants.CONFIG.Namespace && ok {
		panic("Invalid request. Namespace not found!")
	}
}

func checkInternalMethod(request structures.Request) {
	if methods.List[request.Method] == nil {
		panic("Invalid request. Method not found!")
	}

}

func checkExternalMethod(request structures.Request) {
	methodSettings, ok := constants.InfrastructureData.Infrastructure[request.Namespace].Methods[request.Method]

	if !ok || (ok && constants.CONFIG.UseIsInternal == true && methodSettings.IsInternal == true) {
		panic("Invalid request. Method not found!")
	}
}

func checkToken(request structures.Request) {
}

func sendCachedResponse(request structures.Request) {

}

func cacheResponse(message map[string] interface{}) {
	namespace := message["namespace"].(string)
	method := message["method"].(string)
	cacheKey := message["cacheKey"].(string)

	methodSettings := constants.InfrastructureData.Infrastructure[namespace].Methods[method]

	seconds := methodSettings.Cache / 1000
	result := message["result"].(map[string] interface{})

	err := entities.Redis.Client.Set(cacheKey, result, time.Duration(seconds)).Err()
	FailOnError(err, "Error on cache response in redis.")
	if err == nil {
		log.Printf("%s Response with namespace = %s, method = %s was cached!", constants.HEADER_REDIS_MESSAGE, namespace, method)
	}
}

func getCacheKey(request structures.Request) string {
	return "cache key"
}

func getDeliveryKey(request structures.Request) string {
	return "delivery key"
}

func sendResponseToClient(parsedMessage map[string]interface{}, fromCache bool) {
	source := parsedMessage["source"].(string)
	deliveryKey := parsedMessage["deliveryKey"]

	if constants.CONFIG.UseCache == true && parsedMessage["cacheKey"] != nil && fromCache == false {
		cacheResponse(parsedMessage)
	}

	switch true {
	case deliveryKey != nil:
		massSending(parsedMessage)
		return
	case source == "http":
		sendByHttp(parsedMessage)
		return
	case source == "ws":
		sendByWs(parsedMessage)
		return
	default:
		log.Printf("%s Unknown source, can't send response %s", constants.HEADER_RMQ_MESSAGE, parsedMessage)
	}
}

func sendByHttp(message map[string]interface{}) {
	channel := make(chan string)
	go func() {
		<-channel
	}()
}

func sendByWs(message map[string]interface{}) {

}

func massSending(message map[string]interface{}) {

}

func applyAfterMiddlewares(parsedMessage map[string] interface{}) {

}

func applyBeforeMiddlewares(request structures.Request) {

}