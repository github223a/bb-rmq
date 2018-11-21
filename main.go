package main

import(
	"./src"
)

func main() {
	go src.RedisInit()
	go src.HttpServerInit()
	src.RmqInit()
}