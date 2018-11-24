package src

import (
	"./constants"
	"./entities"
	"./methods"
	"./structures"
	"fmt"
	"log"
	"net/http"
	"time"
)

func processingExternalMethod(request structures.Request, transport http.ResponseWriter) {
	//fmt.Printf("%+v\n", request)
	cacheTimer := getMethodSettings(request).Cache
	fmt.Println(1)
	go func() {
		enableResponseListener(transport)
	}()
	fmt.Println(2)

	if constants.CONFIG.UseCache == true && cacheTimer > 0 {
		fmt.Println(4)
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
	methodSettings, isExist := namespaceSettings.Methods[request.Method]

	if !isExist || (isExist && constants.CONFIG.UseIsInternal == true && methodSettings.IsInternal == true) {
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
	fmt.Println("message 777 = ", parsedMessage)

	if constants.CONFIG.UseCache == true && parsedMessage["cacheKey"] != nil && fromCache == false {
		cacheResponse(parsedMessage)
	}
	//fmt.Println("message 777 = ", parsedMessage)
	switch true {
	case deliveryKey != nil:
		massSending(parsedMessage)
		break
	case source == "http":
		sendByHttp(parsedMessage)
		break
	case source == "ws":
		sendByWs(parsedMessage)
		break
	default:
		log.Printf("%s Unknown source, can't send response %s", constants.HEADER_RMQ_MESSAGE, parsedMessage)
	}
}

func sendByHttp(message map[string]interface{}) {
	fmt.Println(7)

	fmt.Printf("SEND BY HTTP message = %s", message)
	ch := entities.Emitter.Channels["1"]
	fmt.Printf("\n channel = %v", ch)
	ch <- message
	//channel := make(chan string)
	//go func() {
	//	<-channel
	//}()
}

func sendByWs(message map[string]interface{}) {

}

func massSending(message map[string]interface{}) {

}

func applyAfterMiddlewares(parsedMessage map[string] interface{}) {

}

func applyBeforeMiddlewares(request structures.Request) {

}