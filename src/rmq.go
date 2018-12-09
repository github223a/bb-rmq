package src

import (
	core "bb-core"
	rmq "bb-rmq"
	"fmt"
	"reflect"

	"./entities"
	"./structures"
)

var Rabbit = rmq.RabbitMQ{}

// func getConfigValue(reflectConnection reflect.Type, variable *string, name string) {
// 	if *variable == "" {
// 		field, _ := reflectConnection.FieldByName(name)
// 		value := field.Tag.Get("default")
// 		*variable = value
// 	}
// }

// func getConfigIntValue(reflectConnection reflect.Type, variable *int, name string) {
// 	if *variable == 0 {
// 		field, _ := reflectConnection.FieldByName(name)
// 		value := field.Tag.Get("default")
// 		i64, _ := strconv.ParseInt(value, 10, 32)
// 		*variable = int(i64)
// 	}
// }

// func getQueueOption(queueOptions map[string]interface{}, name string) bool {
// 	if queueOptions[name] == nil {
// 		switch name {
// 		case "durable":
// 			return true
// 		case "autoDelete":
// 			return false
// 		case "noAck":
// 			return false
// 		default:
// 			return true
// 		}
// 	}
// 	return queueOptions[name].(bool)
// }

func getRabbitUrl() string {
	template := "%s://%s:%s@%s:%d"
	protocol, hostname, username, password, port :=
		core.Config.RabbitMQ.Connection.Protocol,
		core.Config.RabbitMQ.Connection.Hostname,
		core.Config.RabbitMQ.Connection.Username,
		core.Config.RabbitMQ.Connection.Password,
		core.Config.RabbitMQ.Connection.Port

	reflectConnection := reflect.TypeOf(CONFIG.RabbitMQ.Connection)

	// getConfigValue(reflectConnection, &protocol, "Protocol")
	// getConfigValue(reflectConnection, &hostname, "Hostname")
	// getConfigValue(reflectConnection, &username, "Username")
	// getConfigValue(reflectConnection, &password, "Password")
	// getConfigIntValue(reflectConnection, &port, "Port")

	return fmt.Sprintf(template, protocol, username, password, hostname, port)
}

func RmqInit() {
	entities.Emitter = entities.CreateEmitter()
	url := getRabbitUrl()
	fmt.Printf("url = %s\n", url)
	core.Rabbit.InitConnection(url)
	core.Rabbit.InitChannels(core.Config.RabbitMQ.Channels)

	forever := make(chan bool)
	core.methods["friendship"].Run(structures.Request{})
	<-forever
}
