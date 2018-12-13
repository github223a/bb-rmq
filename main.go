package main

import (
	core "bb_core"
	"fmt"

	src "./src"
)

func main() {
	core.Data.InitCore()
	x := src.GetConfig()

	fmt.Print(x)
	// rmq.Rabbit.InitConnection(rmq.GetRabbitUrl(src.GetConfig().RabbitMQ))
	// go src.RedisInit()
	// go src.HttpServerInit()
}
