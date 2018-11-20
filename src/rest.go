package src

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
const headerMessage = "HttpServer: "

func getUrl() string {
	template := "http://%s:%d%s"
	host, path, port := Config.Location.Rest.Host, Config.Location.Rest.Path, Config.Location.Rest.Port

	return fmt.Sprintf(template, host, port, path)
}

type Request struct {
	Id   string  `json:"id"`
}

func requestProcessing(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var request Request
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	id := request.Id
	fmt.Println(id)
}

func HttpServerInit() {
	http.HandleFunc(Config.Location.Rest.Path, requestProcessing) // set router
	log.Printf(headerMessage + "Server is starting by url %s", getUrl())

	err := http.ListenAndServe(fmt.Sprintf(":%d", Config.Location.Rest.Port), nil) // set listen port
	if err != nil {
		log.Fatal(headerMessage, err)
	}
}