package main

import (
	core "bb_core"
)

func main() {
	core.Data.Config.Init("./config.development.json")
	// fmt.Printf("%+v\n", core.Data.Config)

	// src.RmqInit()
	// go src.RedisInit()
	// go src.HttpServerInit()
}
