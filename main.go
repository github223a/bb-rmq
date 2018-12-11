package main

import (
	core "bb_core"
)

/*Config lalala*/
var Config ConfigStrucuture

func main() {
	core.Data.InitConfig()
	// fmt.Printf("%+v\n", core.Data.Config)

	// bytes, _ := json.Marshal(core.Data.Config)
	// json.Unmarshal([]byte(bytes), &Config)
	// core.Data.Config = Config

	// fmt.Printf("Config ===== %+v\n", Config)
	// fmt.Printf("ALLALA \n%s", core.Data.Config.Namespace)

	// data, _ := json.MarshalIndent(core.Data.Config, ",", " ")
	// fmt.Printf("%+v\n", string(data))

	// core.Data.Rabbit.InitConnection()
	// go src.RedisInit()
	// go src.HttpServerInit()
}
