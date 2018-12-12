package src

import (
	core "bb_core"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"./entities"
)

func processingExternalMethod(request core.Request, transport http.ResponseWriter) {
	//fmt.Printf("%+v\n", request)
	cacheTimer := getMethodSettings(request).Cache

	go func() {
		enableResponseListener(transport)
	}()

	if GetConfig().UseCache == true && cacheTimer > 0 {
		//request.cacheKey = getCacheKey(request)
		sendCachedResponse(request)
		return
	}
	applyBeforeMiddlewares(request)
	// rmq.sendToInternal(request)
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
}

func processingInternalMethod(request core.Request) {
	validateRequest(request)
	checkNamespace(request)
	checkInternalMethod(request)
	checkToken(request)

	// core.methods[request.Method].Run(request)
}

func validateRequest(request core.Request) {
}

func checkNamespace(request core.Request) {
	namespace := request.Namespace
	_, ok := core.Infrastructure.Data[namespace]

	if namespace != core.NAMESPACE_INTERNAL && !ok {
		panic("Invalid request. Namespace not found!")
	}
}

func checkInternalMethod(request core.Request) {
	// if core.methods[request.Method] == nil {
	// 	panic("Invalid request. Method not found!")
	// }

}

func checkExternalMethod(request core.Request) {
	namespaceSettings := getNamespaceSettings(request)
	methodSettings, isExist := namespaceSettings.Methods[request.Method]

	if !isExist || (isExist && GetConfig().UseIsInternal == true && methodSettings.IsInternal == true) {
		panic("Invalid request. Method not found!")
	}
}

func checkToken(request core.Request) {
}

func sendCachedResponse(request core.Request) {

}

func cacheResponse(response core.SuccessResponse) {
	// methodSettings := getMethodSettings(core.Request{Namespace: response.Namespace, Method: response.Method})
	// seconds := methodSettings.Cache / 1000

	// err := entities.Redis.Client.Set(*response.CacheKey, response.Result, time.Duration(seconds)).Err()
	// FailOnError(err, "Error on cache response.", "redis")
	// if err == nil {
	// 	log.Printf("%s Response with namespace = %s, method = %s was cached!", HEADER_REDIS_MESSAGE, response.Namespace, response.Method)
	// }
}

func getCacheKey(request core.Request) string {
	return "cache key"
}

func sendResponseToClient(response map[string]interface{}, fromCache bool) {
	source := response["source"].(string)

	if GetConfig().UseCache == true && response["cacheKey"] != nil && fromCache == false {
		var successResponse core.SuccessResponse
		respB, _ := json.Marshal(response)
		json.Unmarshal(respB, &successResponse)
		cacheResponse(successResponse)
	}
	//fmt.Println("message 777 = ", parsedMessage)
	switch true {
	case source == "http":
		sendByHttp(response)
		break
	case source == "ws":
		sendByWs(response)
		break
	default:
		log.Printf("%s Unknown source, can't send response %s", core.HEADER_APPLICATION_MESSAGE, response)
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

func applyBeforeMiddlewares(request core.Request) {

}
