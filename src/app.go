package src

import (
	"./constants"
	"./entities"
	methods "./methods"
	"./templates"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func processingExternalMethod(request templates.Request) {
	fmt.Printf("%+v\n", request)
	infrastructure := constants.INFRASTRUCTURE["infrastructure"].(map[string] interface{})
	serviceSettings := infrastructure[request.Namespace].(map[string] interface{})
	serviceMethods := serviceSettings["methods"].(map[string] interface{})
	methodSettings := serviceMethods[request.Method].(map[string] interface{})
	cacheTimer := methodSettings["cache"].(float64)

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
	log.Printf("%s Sent message to [* %s *]. Message %s", constants.HEADER_RMQ_MESSAGE, constants.NAMESPACE_INTERNAL, _request)
}

func processingInternalMethod(message templates.Request) {
	validateRequest(message)
	checkNamespace(message)
	checkInternalMethod(message)
	checkToken(message)
	runMethod(message)
}

func validateRequest(request templates.Request) {
}
func checkNamespace(request templates.Request) {
}
func checkInternalMethod(request templates.Request) {
}
func checkExternalMethod(request templates.Request) {
}
func checkToken(request templates.Request) {
}

func runMethod(request templates.Request) {
	method := methods.List[request.Method]

	if method == nil {
		fmt.Println("no method")
		return // need send error to client
	}
	method.Run(request)
}

func sendCachedResponse(request templates.Request) {

}

func cacheResponse(message map[string] interface{}) {
	namespace := message["namespace"].(string)
	method := message["method"].(string)
	cacheKey := message["cacheKey"].(string)

	infrastructure := constants.INFRASTRUCTURE["infrastructure"].(map[string] interface{})
	serviceSettings := infrastructure[namespace].(map[string] interface{})
	serviceMethods := serviceSettings["methods"].(map[string] interface{})
	methodSettings := serviceMethods[method].(map[string] interface{})

	cacheTimer := methodSettings["cache"].(int)
	seconds := cacheTimer / 1000
	result := message["result"].(map[string] interface{})

	err := entities.Redis.Client.Set(cacheKey, result, time.Duration(seconds)).Err()
	FailOnError(err, "Error on cache response in redis.")
	if err == nil {
		log.Printf("%s Response with namespace = %s, method = %s was cached!", constants.HEADER_REDIS_MESSAGE, namespace, method)
	}
}

func getCacheKey(request templates.Request) string {
	return "cache key"
}

func getDeliveryKey(request templates.Request) string {
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

func applyBeforeMiddlewares(request templates.Request) {

}