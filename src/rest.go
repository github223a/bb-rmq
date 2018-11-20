package src

import (
	"./constants"
	"./templates"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getUrl() string {
	template := "http://%s:%d%s"
	host, path, port := constants.CONFIG.Location.Rest.Host, constants.CONFIG.Location.Rest.Path, constants.CONFIG.Location.Rest.Port

	return fmt.Sprintf(template, host, port, path)
}

func parseRequest(req *http.Request, variable *templates.Request) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&*variable)
	FailOnError(err, "Error on parse request.")
}

func httpRequestProcessing(writer http.ResponseWriter, req *http.Request) {
	var request templates.Request
	parseRequest(req, &request)
	setSource(&request, "http")
	logRequest(request, "http")
}

func HttpServerInit() {
	http.HandleFunc(constants.CONFIG.Location.Rest.Path, httpRequestProcessing) // set router
	log.Printf(constants.HEADER_HTTP_MESSAGE + "Server is starting by url %s", getUrl())

	err := http.ListenAndServe(fmt.Sprintf(":%d", constants.CONFIG.Location.Rest.Port), nil) // set listen port
	FailOnError(err, "Error on start http server.")
}