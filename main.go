package main

import (
	core "bb_core"
	"fmt"
)

var Config = &ConfigStructure{}

func main() {
	core.Data.InitConfig()
	// fmt.Printf("%+v\n", core.Data.Config)

	// ddd, _ := json.Marshal(core.Data.Config)
	// json.Unmarshal([]byte(ddd), &Config)
	// core.Data.Config = Config

	fmt.Printf("%+v\n", core.Data.Config)

	// data, _ := json.MarshalIndent(core.Data.Config, ",", " ")
	// fmt.Printf("%+v\n", string(data))

	// core.Data.Rabbit.InitConnection()
	// go src.RedisInit()
	// go src.HttpServerInit()
}
