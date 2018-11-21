package src

import (
	"./constants"
	"./entities"
	methods "./internal-methods"
	"./templates"
	"fmt"
	"log"
	"time"
)

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
func checkToken(request templates.Request) {
}
func runMethod(request templates.Request) {
	method := methods.List[request.Method]

	if method == nil {
		fmt.Println("no method")
		return // need send error to client
	}
	method(request)
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
	if err != nil {
		FailOnError(err, "Error on set key in redis.")
	}
	log.Printf("%s Response with namespace = %s, method = %s was cached!", constants.HEADER_REDIS_MESSAGE, namespace, method)
}

func sendResponseToClient(parsedMessage map[string]interface{}, fromCache bool) {
	source := parsedMessage["source"].(string)
	deliveryKey := parsedMessage["deliveryKey"]

	if constants.CONFIG.UseCache == true && parsedMessage["cacheKey"] != nil && fromCache == false {
		cacheResponse(parsedMessage)
	}

	if deliveryKey != nil {
		massSending(parsedMessage)
		return
	}

	if source == "http" {
		sendByHttp(parsedMessage)
		return
	}

	if source == "ws" {
		sendByWs(parsedMessage)
		return
	}
	log.Printf("%s Unknown source, can't send response %s", constants.HEADER_RMQ_MESSAGE, parsedMessage)
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