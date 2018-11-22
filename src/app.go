package src

import (
	"./constants"
	"./entities"
	"./methods"
	"./structures"
	"fmt"
	"log"
	"time"
)

func processingExternalMethod(request structures.Request) {
	//fmt.Printf("%+v\n", request)
	cacheTimer := getMethodSettings(request).Cache

	if constants.CONFIG.UseCache == true && cacheTimer > 0 {
		//request.cacheKey = getCacheKey(request)
		sendCachedResponse(request)
		return
	}
	applyBeforeMiddlewares(request)
	sendToInternal(request)
}

func processingInternalMethod(request structures.Request) {
	validateRequest(request)
	checkNamespace(request)
	checkInternalMethod(request)
	checkToken(request)

	methods.List[request.Method].Run(request)
}

func validateRequest(request structures.Request) {
}

func checkNamespace(request structures.Request) {
	namespace := request.Namespace
	_, ok := constants.InfrastructureData.Infrastructure[namespace]

	if namespace != constants.NAMESPACE_INTERNAL && !ok {
		panic("Invalid request. Namespace not found!")
	}
}

func checkInternalMethod(request structures.Request) {
	if methods.List[request.Method] == nil {
		panic("Invalid request. Method not found!")
	}

}

func checkExternalMethod(request structures.Request) {
	namespaceSettings := getNamespaceSettings(request)
	methodSettings, isExist := namespaceSettings["methods"].(map[string] interface{})[request.Method]
	fmt.Printf("req = %+v\n",  request)
	fmt.Printf("lala %v %s", isExist, methodSettings)
	methodSettings2 := methodSettings.(map[string] interface{})
	if !isExist || (isExist && constants.CONFIG.UseIsInternal == true && methodSettings2["isInternal"] == true) {
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
	methodSettings := getMethodSettings(structures.Request{Namespace:namespace, Method:method})

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