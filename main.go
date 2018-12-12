package main

import (
	core "bb_core"
	"fmt"

	src "./src"
)

func main() {
	core.Data.InitCore()

	x := src.GetConfig()
	fmt.Printf("%+v\n", x.Namespace)

	// fmt.Printf("%+v\n", core.Data.Infrastructure)
	// rmq.Rabbit.InitConnection(rmq.GetRabbitUrl()
	// go src.RedisInit()
	// go src.HttpServerInit()
}
