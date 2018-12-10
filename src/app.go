package src

import (
	core "bb_core"
	rmq "bb_rmq"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"./entities"
)

func processingExternalMethod(request rmq.Request, transport http.ResponseWriter) {
	//fmt.Printf("%+v\n", request)
	cacheTimer := getMethodSettings(request).Cache

	go func() {
		enableResponseListener(transport)
	}()

	if CONFIG.UseCache == true && cacheTimer > 0 {
		//request.cacheKey = getCacheKey(request)
		sendCachedResponse(request)
		return
	}
	applyBeforeMiddlewares(request)
	rmq.sendToInternal(request)
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
}

func processingInternalMethod(request rmq.Request) {
	validateRequest(request)
	checkNamespace(request)
	checkInternalMethod(request)
	checkToken(request)

	core.methods[request.Method].Run(request)
}

func validateRequest(request rmq.Request) {
}

func checkNamespace(request rmq.Request) {
	namespace := request.Namespace
	_, ok := InfrastructureData.Infrastructure[namespace]

	if namespace != NAMESPACE_INTERNAL && !ok {
		panic("Invalid request. Namespace not found!")
	}
}

func checkInternalMethod(request rmq.Request) {
	if core.methods[request.Method] == nil {
		panic("Invalid request. Method not found!")
	}

}

func checkExternalMethod(request rmq.Request) {
	namespaceSettings := getNamespaceSettings(request)
	methodSettings, isExist := namespaceSettings.Methods[request.Method]

	if !isExist || (isExist && CONFIG.UseIsInternal == true && methodSettings.IsInternal == true) {
		panic("Invalid request. Method not found!")
	}
}

func checkToken(request rmq.Request) {
}

func sendCachedResponse(request rmq.Request) {

}

func cacheResponse(response rmq.SuccessResponse) {
	methodSettings := getMethodSettings(rmq.Request{Namespace: response.Namespace, Method: response.Method})
	seconds := methodSettings.Cache / 1000

	err := entities.Redis.Client.Set(*response.CacheKey, response.Result, time.Duration(seconds)).Err()
	FailOnError(err, "Error on cache response.", "redis")
	if err == nil {
		log.Printf("%s Response with namespace = %s, method = %s was cached!", HEADER_REDIS_MESSAGE, response.Namespace, response.Method)
	}
}

func getCacheKey(request rmq.Request) string {
	return "cache key"
}

func sendResponseToClient(parsedMessage map[string]interface{}, fromCache bool) {
	source := parsedMessage["source"].(string)
	deliveryKey := parsedMessage["deliveryKey"]

	if CONFIG.UseCache == true && parsedMessage["cacheKey"] != nil && fromCache == false {
		var successResponse rmq.SuccessResponse
		respB, _ := json.Marshal(parsedMessage)
		json.Unmarshal(respB, &successResponse)
		cacheResponse(successResponse)
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
		log.Printf("%s Unknown source, can't send response %s", HEADER_RMQ_MESSAGE, parsedMessage)
	}
}

func sendByHttp(message map[string]interface{}) {
	ch := entities.Emitter.Channels["1"]
	ch <- message
}

func sendByWs(message map[string]interface{}) {

}

func massSending(message map[string]interface{}) {

}

func applyAfterMiddlewares(parsedMessage map[string]interface{}) {

}

func applyBeforeMiddlewares(request rmq.Request) {

}
