package src

import (
	core "bb_core"
	"fmt"
)

func GetConfig() ConfigStructure {
	fmt.Printf("\11111111 %+v\n", core.Data.Config)

	config, ok := core.Data.Config.(ConfigStructure)
	fmt.Printf("\nconfig %+v\n", config)

	if !ok {
		panic("Can't structuring config.")
	}

	// var config ConfigStructure

	// bytes, err := json.Marshal(core.Data.Config)

	// if err != nil {
	// 	log.Fatal("Error on marshal config", err)
	// }

	// err2 := json.Unmarshal(bytes, &config)
	// if err2 != nil {
	// 	log.Fatal("error on unmarshal config", err2)
	// }

	return config
}

func getNamespaceSettings(request core.Request) core.InfrastructureData {
	return core.Infrastructure.Data[request.Namespace]
}

func getMethodSettings(request core.Request) core.MethodSettings {
	return getNamespaceSettings(request).Methods[request.Method]
}

// func logRequest(request core.Request, t string) {
// 	fmt.Printf("%s Get request %+v\n", getMessageHeader(t), request)
// }

func setSource(request *core.Request, value string) {
	request.Source = value
}

type ConfigStructure struct {
	Namespace     string     `json:"namespace"`
	UseCache      bool       `json:"useCache"`
	UseIsInternal bool       `json:"useIsInternal"`
	Location      Location   `json:"location"`
	Redis         core.Redis `json:"redis"`
	core.CommonConfig
	// Log      core.Log      `json:"log"`
	// RabbitMQ core.RabbitMQ `json:"rabbitMQ"`
	// PingTimeout int           `json:"pingTimeout"`
}

type Location struct {
	Ws   Ws   `json:"ws"`
	Rest Rest `json:"rest"`
}

type Ws struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
	Path string `json:"path"`
}

type Rest struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
	Path string `json:"path"`
}

// {
// 	Namespace:
// 	UseCache:false
// 	UseIsInternal:false
// 	Location:{
// 		Ws:{
// 			Host: Port:0 Path:
// 			}
// 		Rest:{Host: Port:0 Path:}
// 	}
// 	Redis:{Host: Port:0 Password:}
// 	Log:{Type: FilePath: FileLevel: ConsoleLevel:}
// 	RabbitMQ:{Connection:{Protocol: Hostname: Username: Password: Port:0} Channels:map[] TestChannelName: TestChannelPingTimeout:0 InfrastructureBindingKey:}
// 	PingTimeout:0}
