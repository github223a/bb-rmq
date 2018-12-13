package main

import (
	core "bb_core"
	rmq "bb_rmq"

	src "./src"
)

func main() {
	core.Data.InitCore()
	rabbitConfig := src.GetConfig().RabbitMQ

	rmq.Rabbit.InitConnection(rmq.GetRabbitUrl(rabbitConfig))
	go rmq.Rabbit.InitChannels(rabbitConfig.Channels)
	go src.HttpServerInit()

	// core.methods["friendship"].Run(Rabbit, rmq.Request{})

	forever := make(chan bool)
	<-forever

	// src.RedisInit()
	// x := src.GetConfig()
	// fmt.Print(x)

}

// type ExampleBaseConfigEntry struct {
// 	Field1 string `json:"field_1"`
// 	Field2 int    `json:"field_2"`
// }

// type ExampleBaseConfig struct { // common
// 	Entry ExampleBaseConfigEntry `json:"entry"`
// }

// type ExampleDerivedConfig struct { // service
// 	ExampleBaseConfig
// 	Field3 float32 `json:"field_3"`
// }

// func getDerivedConfig() interface{} {
// 	return &ExampleDerivedConfig{}
// }

// type ExampleConfigContainer struct { //core
// 	Config interface{}
// }

// func initConfigContainer(jsonSrc string, derivedConfFunc func() interface{}) *ExampleConfigContainer {
// 	conf := derivedConfFunc()
// 	err := json.Unmarshal([]byte(jsonSrc), conf)

// 	if err != nil {
// 		panic(err)
// 	}

// 	container := ExampleConfigContainer{Config: conf}

// 	return &container
// }

// func getConfig(container *ExampleConfigContainer) *ExampleDerivedConfig {
// 	conf, ok := container.Config.(*ExampleDerivedConfig)
// 	if !ok {
// 		panic("Can't cast config")
// 	}

// 	return conf
// }

// func main() {
// 	jsonConf := `{"entry": {"field_1": "testField", "field_2": 69}, "field_3": 13.37}`

// 	cont := initConfigContainer(jsonConf, getDerivedConfig)

// 	conf := getConfig(cont)

// 	fmt.Println(conf.Field3)
// }
