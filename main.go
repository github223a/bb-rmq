package main

import (
	core "bb_core"
)

type Aaa struct {
	A string `json:"a,omitempty"`
	B int    `json:"b,omitempty"`
}

var ss = Aaa{B: 1}

func main() {
	core.Data.Config.Init("./config.development.json")
	// fmt.Printf("%+v\n", core.Data.Config)
	// fmt.Printf("%+v", ss)

	// src.RmqInit()
	// go src.RedisInit()
	// go src.HttpServerInit()
}
