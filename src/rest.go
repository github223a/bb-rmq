package src

import (
	"./constants"
	"./structures"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HttpServerInit() {
	http.HandleFunc(constants.CONFIG.Location.Rest.Path, httpRequestProcessing) // set router
	log.Printf(constants.HEADER_HTTP_MESSAGE + "Server is starting by url %s", getUrl())

	err := http.ListenAndServe(fmt.Sprintf(":%d", constants.CONFIG.Location.Rest.Port), nil) // set listen port
	FailOnError(err, "Error on start http server.")
}

func httpRequestProcessing(writer http.ResponseWriter, req *http.Request) {
	var request structures.Request
	parseRequest(req, &request)
	setSource(&request, "http")

	validateRequest(request)
	checkNamespace(request)
	checkExternalMethod(request)
	processingExternalMethod(request)
	logRequest(request, "http")
}

func getUrl() string {
	template := "http://%s:%d%s"
	host, path, port := constants.CONFIG.Location.Rest.Host, constants.CONFIG.Location.Rest.Path, constants.CONFIG.Location.Rest.Port

	return fmt.Sprintf(template, host, port, path)
}

func parseRequest(req *http.Request, variable *structures.Request) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&*variable)
	FailOnError(err, "Error on parse request.")
}
