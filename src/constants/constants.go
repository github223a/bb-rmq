package constants

import (
	"../templates"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func FailOnErrorReadConfig(err error, msg string) {
	if err != nil {
		log.Panicf("%s", msg)
		panic(err)
	}
}

func readConfig() templates.Config {
	var config templates.Config

	configFile, err := os.Open("./config.development.json")
	FailOnErrorReadConfig(err, "Error on open config file.")
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal([]byte(byteValue), &config)
	return config
}

var CONFIG = readConfig()
var INFRASTRUCTURE = map[string] interface{}{}

const NAMESPACE_INTERNAL = "internal"
const HEADER_RMQ_MESSAGE = "[*] RabbitMQ: "
const HEADER_HTTP_MESSAGE = "[*] HttpServer: "
const HEADER_WS_MESSAGE = "[*] WsServer: "
const HEADER_REDIS_MESSAGE = "[*] Redis: "
const HEADER_UNKNOWN = "[*] Unknown: "
