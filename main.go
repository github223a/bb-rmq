package main

import (
	"./src"
)

func main() {
	src.RmqInit()
	go src.RedisInit()
	go src.HttpServerInit()
}
