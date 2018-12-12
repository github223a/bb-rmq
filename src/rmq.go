package src

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

func RmqInit() {
	// entities.Emitter = entities.CreateEmitter()
	// url := getRabbitUrl()
	// fmt.Printf("url = %s\n", url)

	// rabbit := Rabbit.InitConnection(url)
	// rabbit.InitChannels(conf.RabbitMQ.Channels)

	// forever := make(chan bool)
	// core.methods["friendship"].Run(Rabbit, rmq.Request{})

	// <-forever
}

type connect interface {
	init()
	close()
}
