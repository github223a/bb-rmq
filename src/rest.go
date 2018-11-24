package src

import (
	"./constants"
	"./entities"
	"./structures"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HttpServerInit() {
	//var (
	//	// flagPort is the open port the application listens on
	//	flagPort = flag.String("port", "7777", "Port to listen on")
	//)
	//mux := http.NewServeMux()
	//mux.HandleFunc(constants.CONFIG.Location.Rest.Path, httpPostRequestHandler)
	//log.Printf("listening on port %s", *flagPort)
	//log.Fatal(http.ListenAndServe(":" + string(constants.CONFIG.Location.Rest.Port), mux))

	http.HandleFunc(constants.CONFIG.Location.Rest.Path, httpPostRequestHandler) // set router
	log.Printf(constants.HEADER_HTTP_MESSAGE + "Server is starting by url %s", getUrl())
	err := http.ListenAndServe(fmt.Sprintf(":%d", constants.CONFIG.Location.Rest.Port), nil) // set listen port
	FailOnError(err, "Error on start http server.")
}

func httpPostRequestHandler(writer http.ResponseWriter, req *http.Request) {
	var request structures.Request

	parseRequest(req, &request)
	setSource(&request, "http")
	logRequest(request, "http")

	validateRequest(request)
	checkNamespace(request)
	checkExternalMethod(request)
	processingExternalMethod(request, writer)
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

func enableResponseListener(transport http.ResponseWriter) {
	channel := make(chan interface{})
	entities.Emitter.Channels["1"] = channel
	msg := <- entities.Emitter.Channels["1"]
	defer close(entities.Emitter.Channels["1"])

	fmt.Println("\n msg = ", msg)
	response, err := json.Marshal(msg)

	if err != nil {
		http.Error(transport, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(3)

	transport.Header().Set("Content-Type", "application/json")
	transport.Write(response)
}