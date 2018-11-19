package rmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"os"
)

const GUEST =  "guest"

var CONFIG = readConfig()
var rabbitMQ = CONFIG["rabbitMQ"].(map[string] interface{})

func readConfig() map[string]interface{} {
	var config map[string]interface{}

	configFile, err := os.Open("./src/config.development.json")
	if err != nil {
		fmt.Println(err)
	}

	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal([]byte(byteValue), &config)

	return config
}

func getRabbitUrl() string{
	var url string

	connection := rabbitMQ["connection"].(map[string] interface{})
	protocol := connection["protocol"].(string)
	hostname := connection["hostname"].(string)
	username, userOk := connection["username"].(string)
	password, passOk := connection["password"].(string)
	port, portOk := connection["port2"].(string)

	url = fmt.Sprintf("%s://", protocol)

	if userOk && passOk {
		url += username + ":" + password
	} else {
		url += GUEST + ":" + GUEST
	}

	url += "@" + hostname + ":"

	if portOk {
		url += string(port)
	} else {
		url += "5672"
	}

	return url
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
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()


	channels := rabbitMQ["channels"].(map[string] interface{})
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

	//fmt.Println("json = ", requestMsg)

	//err = ch.Publish(
	//	"",      // exchange
	//	q.Name,           // routing key
	//	false,  // mandatory
	//	false,  // immediate
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte(requestMsg),
	//	})
	//log.Printf(" [x] Sent %s", requestMsg)
	//failOnError(err, "Failed to publish a message")
}