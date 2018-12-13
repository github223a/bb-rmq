package src

import (
	core "bb_core"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"./entities"
)

func HttpServerInit() {
	rest := GetConfig().Location.Rest
	http.HandleFunc(rest.Path, postHandler) // set router
	port := strconv.FormatInt(rest.Port, 10)
	log.Printf(HEADER_HTTP_MESSAGE+"Server is starting by url %s", getUrl())
	err := http.ListenAndServe(port, nil) // set listen port
	core.FailOnError(HEADER_HTTP_MESSAGE, "Error on start http server.", err)
}

func postHandler(writer http.ResponseWriter, req *http.Request) {
	var request core.Request

	parseRequest(req, writer, &request)
	setSource(&request, "http")
	// logRequest(request, "http")
	validateRequest(request)
	checkNamespace(request)
	checkExternalMethod(request)
	processingExternalMethod(request, writer)
}

func getUrl() string {
	template := "http://%s:%d%s"
	host, path, port := GetConfig().Location.Rest.Host, GetConfig().Location.Rest.Path, GetConfig().Location.Rest.Port

	return fmt.Sprintf(template, host, port, path)
}

func parseRequest(req *http.Request, writer http.ResponseWriter, variable *core.Request) {
	if req.Method != "POST" {
		http.Error(writer, http.StatusText(405), 405)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&*variable)
	core.FailOnError(core.HEADER_APPLICATION_MESSAGE, "Error on parse request.", err)
}

func enableResponseListener(transport http.ResponseWriter) {
	entities.Emitter.Channels["1"] = make(chan interface{})
	response := <-entities.Emitter.Channels["1"]
	defer close(entities.Emitter.Channels["1"])
	defer delete(entities.Emitter.Channels, "1")

	responseB, _ := json.Marshal(response)
	//transport.Header().Del("Content-Length")
	transport.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//transport.Header().Set("Content-Length", strconv.Itoa(len(responseB)))
	_, writeErr := transport.Write(responseB)
	fmt.Println("write error", writeErr)
	log.Printf("%s Response %s was sent.", HEADER_HTTP_MESSAGE, response)
}
