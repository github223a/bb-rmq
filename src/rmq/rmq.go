package rmq

import (
	"../templates"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

const GUEST =  "guest"
const RABBIT_PORT = "5672"

var config = readConfig()

func readConfig() templates.Config {
	var config templates.Config

	configFile, err := os.Open("./src/config.development.json")
	if err != nil {
		fmt.Println(err)
	}

	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal([]byte(byteValue), &config)
	fmt.Println(config)
	return config
}

func getUserAndPassword(username, password string) string {
	//result = username + ":" + password
	//result = GUEST + ":" + GUEST
	return username + ":" + password
}

//func getPort(port string) string {
//	var result string
//
//	if portOk {
//		result = string(port)
//	} else {
//		result = RABBIT_PORT
//	}
//	return result
//}

func getRabbitUrl() string {
	var url string
	//var protocol, hostname, username, password string
	//var port int
	var keys = [5]string {
		"Protocol",
		"Hostname",
		"Username",
		"Password",
		"Port",
	}
	var reflectConnection = reflect.TypeOf(config.Connection)

	for i := 0; i < reflectConnection.NumField(); i++ {
		fmt.Println(i)
	}
	url = "%s://%s@%s:%s"

	for _, name := range keys {
		field, found := reflectConnection.FieldByName(name)
		fmt.Println("field = ", field)
		fmt.Println("found = ", found)
		fmt.Println("value = ", field)

		if found {
			url = fmt.Sprintf(field.Tag.Get(name))
		}
	}

	fmt.Println("url 1 = ", url)

	//defaultHostname := field.Tag.Get("default")
	//fmt.Println("field", field.Tag.Get("default"), found)


	return url
	//return fmt.Sprintf("%s://%s@%s:%s",
	//	protocol,
	//	getUserAndPassword(username, password),
	//	hostname,
	//	string(port),
	//)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func declareQueue (ch *amqp.Channel, settings map[string] interface{}) {
	queueName := settings["queueName"].(string)
	queueOptions := settings["queueOptions"].(map[string] interface{})

	args := make(amqp.Table)
	args["x-message-ttl"] = int32(30000)

	_, queueError := ch.QueueDeclare(
		queueName, // name
		queueOptions["durable"].(bool),   // durable
		queueOptions["autoDelete"].(bool),   // delete when unused
		false,   // exclusive
		false,   // no-wait
		args,     // arguments
	)
	failOnError(queueError, "Failed to declare a queue")
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
	failOnError(err, "Failed to declare an exchange")
}

func declareCunsumer (ch *amqp.Channel, settings map[string] interface{}) {
	queueName := settings["queueName"].(string)
	queueOptions := settings["queueOptions"].(map[string] interface{})
	msgs, err := ch.Consume(
		queueName, // queue
		"",     // consumer
		queueOptions["noAck"].(bool),   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message from %s: %s", queueName, d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages from %s. To exit press CTRL+C", queueName)
}

func Init() {
	url := getRabbitUrl()
	fmt.Println(url)
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to rabbitMQ")
	defer conn.Close()


	channels := config.Channels
	forever := make(chan bool)

	for key, _ := range channels {
		channel, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer channel.Close()

		fmt.Println(key)
		settings := channels[key].(map[string] interface{})
		consumeActivate := settings["consumeActivate"].(bool)

		declareExchange(channel, settings)
		declareQueue(channel, settings)

		if consumeActivate {
			declareCunsumer(channel, settings)
		}
	}
	<-forever

	//request := templates.Handshake()
	//jsonData, err := json.Marshal(request)
	//requestMsg := string(jsonData)
	//
	//fmt.Println("json = ", requestMsg)
	//
	//err = channel.Publish(
	//	"",      // exchange
	//	queue.Name,           // routing key
	//	false,  // mandatory
	//	false,  // immediate
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte(requestMsg),
	//	})
	//log.Printf(" [x] Sent %s", requestMsg)
	//failOnError(err, "Failed to publish a message")
}