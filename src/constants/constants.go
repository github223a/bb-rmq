package constants

import (
	"../structures"
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

func readConfig() structures.Config {
	var config structures.Config

	configFile, err := os.Open("./config.development.json")
	FailOnErrorReadConfig(err, "Error on open config file.")
	defer configFile.Close()

	//decoder := json.NewDecoder(configFile)
	//decodeErr := decoder.Decode(&config)
	//FailOnErrorReadConfig(decodeErr, "Error on decode config.")

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal([]byte(byteValue), &config)
	return config
}

const NAMESPACE_INTERNAL = "internal"
const HEADER_RMQ_MESSAGE = "[*] RabbitMQ: "
const HEADER_HTTP_MESSAGE = "[*] HttpServer: "
const HEADER_WS_MESSAGE = "[*] WsServer: "
const HEADER_REDIS_MESSAGE = "[*] Redis: "
const HEADER_UNKNOWN = "[*] Unknown: "

var CONFIG = readConfig()
var InfrastructureData structures.InfrastructureData
