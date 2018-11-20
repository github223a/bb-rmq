package main

import(
	"./src"
)

func main() {
	go src.HttpServerInit()
	src.RmqInit()
}