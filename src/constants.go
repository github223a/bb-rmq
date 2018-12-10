package src

import (
	core "bb_core"
)

func cacheConfig() core.Config {
	conf := core.Config{}
	core.Config.Init("./config.development.json")

	// data, _ := json.MarshalIndent(conf, "", " ")
	// fmt.Printf("%s\n", data)
	return conf
}

var CONFIG = cacheConfig()
var InfrastructureData core.InfrastructureData

const NAMESPACE_INTERNAL = "internal"
const HEADER_HTTP_MESSAGE = "[*] HttpServer: "
const HEADER_WS_MESSAGE = "[*] WsServer: "
const HEADER_REDIS_MESSAGE = "[*] Redis: "
const HEADER_UNKNOWN = "[*] Unknown: "
