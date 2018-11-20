package src

import (
	"./constants"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"reflect"
	"strconv"
)

func getConfigValue(reflectConnection reflect.Type, variable *string, name string ) {
	if *variable == "" {
		field, _ := reflectConnection.FieldByName(name)
		value := field.Tag.Get("default")
		*variable = value
	}
}

func getConfigIntValue(reflectConnection reflect.Type, variable *int, name string ) {
	if *variable == 0 {
		field, _ := reflectConnection.FieldByName(name)
		value := field.Tag.Get("default")
		i64, _ := strconv.ParseInt(value, 10, 32)
		*variable = int(i64)
	}
}

func getQueueOption(queueOptions map[string] interface{}, name string) bool {
	if queueOptions[name] == nil {
		switch name {
			case "durable":
				return true
			case "autoDelete":
				return false
			case "noAck":
				return false
			default:
				return true
		}
	}
	return queueOptions[name].(bool)
}

func getRabbitUrl() string {
	template := "%s://%s:%s@%s:%d"
	protocol, hostname, username, password, port :=
		constants.CONFIG.RabbitMQ.Connection.Protocol,
		constants.CONFIG.RabbitMQ.Connection.Hostname,
		constants.CONFIG.RabbitMQ.Connection.Username,
		constants.CONFIG.RabbitMQ.Connection.Password,
		constants.CONFIG.RabbitMQ.Connection.Port

	reflectConnection := reflect.TypeOf(constants.CONFIG.RabbitMQ.Connection)

	getConfigValue(reflectConnection, &protocol, "Protocol")
	getConfigValue(reflectConnection, &hostname, "Hostname")
	getConfigValue(reflectConnection, &username, "Username")
	getConfigValue(reflectConnection, &password, "Password")
	getConfigIntValue(reflectConnection, &port, "Port")

	return fmt.Sprintf(template, protocol, username, password, hostname, port)
}

func declareQueue (ch *amqp.Channel, settings map[string] interface{}) {
	queueName := settings["queueName"].(string)
	queueOptions := settings["queueOptions"].(map[string] interface{})

	args := make(amqp.Table)
	args["x-message-ttl"] = int32(queueOptions["messageTtl"].(float64))

	_, queueError := ch.QueueDeclare(
		queueName, // name
		getQueueOption(queueOptions, "durable"),   // durable
		getQueueOption(queueOptions, "autoDelete"),   // delete when unused
		false,   // exclusive
		false,   // no-wait
		args,     // arguments
	)
	FailOnError(queueError, "Failed to declare a queue")
}

func declareExchange (ch *amqp.Channel, settings map[string] interface{}) {
	exchangeName := settings["exchangeName"].(string)
	exchangeType := settings["exchangeType"].(string)

	err := ch.ExchangeDeclare(
		exchangeName,   // name
		exchangeType, // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	FailOnError(err, "Failed to declare an exchange")
}

func bindQueue(ch *amqp.Channel, settings map[string] interface{}) {
	queueName := settings["queueName"].(string)
	exchangeName := settings["exchangeName"].(string)
	bindingKey := settings["bindingKey"].(string)

	err := ch.QueueBind(
		queueName,      // queue name
		bindingKey,    // routing key
		exchangeName, // exchange
		false,
		nil)
	FailOnError(err, "Failed to bind a queue")
}

func declareCunsumer (ch *amqp.Channel, settings map[string] interface{}) {
	queueName := settings["queueName"].(string)
	queueOptions := settings["queueOptions"].(map[string] interface{})

	msgs, err := ch.Consume(
		queueName, // queue
		"",     // consumer
		getQueueOption(queueOptions, "noAck"),   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for message := range msgs {
			log.Printf("Received a message from %s: %s", queueName, message.Body)
			rmqProcessing(message.Body)
		}
	}()
	if queueName == "" {
		queueName = settings["bindingKey"].(string)
	}
	log.Printf("%s Waiting for messages from %s channel. To exit press CTRL+C", constants.HEADER_RMQ_MESSAGE, queueName)
}

func rmqProcessing(message []byte) {
	var parsedMessage map[string] interface{}
	unMarshalMessage(message, &parsedMessage)

	if parsedMessage["error"] == nil && parsedMessage["result"] == nil {
		processingInternalMethod(parsedMessage)
		return
	}
	//if parsedMessage["result"] != nil {
	//	applyAfterMiddlewares(parsedMessage)
	//}
	//
	//sendResponseToClient(parsedMessage)
}

func RmqInit() {
	url := getRabbitUrl()
	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to connect to rabbitMQ")
	defer conn.Close()

	channels := constants.CONFIG.RabbitMQ.Channels
	forever := make(chan bool)

	for key, _ := range channels {
		channel, err := conn.Channel()
		FailOnError(err, "Failed to open a channel")
		defer channel.Close()

		settings := channels[key].(map[string] interface{})
		consumeActivate := settings["consumeActivate"].(bool)
		bindingKey := settings["bindingKey"]

		declareExchange(channel, settings)
		declareQueue(channel, settings)

		if bindingKey != nil {
			bindQueue(channel, settings)
		}

		if consumeActivate {
			declareCunsumer(channel, settings)
		}
	}
	<-forever
}