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

	url = "%s://%s:%s@%s:%s"

	protocol, hostname, username, password, port :=
		config.Connection.Protocol,
		config.Connection.Hostname,
		config.Connection.Username,
		config.Connection.Password,
		config.Connection.Port
	fmt.Println(protocol)
	fmt.Println(hostname)
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(port)

	if username == "" {
		reflectConnection := reflect.TypeOf(config.Connection)
		field, _ := reflectConnection.FieldByName("Username")
		value := field.Tag.Get("default")
		username = value
		fmt.Println("username 2 = ", username)
	}

	//var keys = [5]string {
	//	"protocol",
	//	"hostname",
	//	"username",
	//	"password",
	//	"port",
	//}

	//for _, name := range keys {
	//	field := config.Connection[name]
	//	fmt.Println(field)
	//}
	//reflectConnection := reflect.TypeOf(config.Connection)

	//for i := 0; i < reflectConnection.NumField(); i++ {
	//	field := reflectConnection.Field(i)
	//	name := field.Tag.Get("json")
	//	fmt.Println("value = ", name)
	//	fmt.Println("field = ", field)
	//
	//}

	//for _, name := range keys {
	//	field, found := reflectConnection.FieldByName(name)
	//
	//	//r := reflect.ValueOf(reflectConnection)
	//	//f := reflect.Indirect(r).FieldByName(name)
	//
	//	//fmt.Println("field = ", field)
	//	//fmt.Println("found = ", found)
	//	//fmt.Println("value = ", f)
	//
	//	//if found {
	//	//	url = fmt.Sprintf(field.Tag.Get(name))
	//	//}
	//}

	//defaultHostname := field.Tag.Get("default")
	//fmt.Println("field", field.Tag.Get("default"), found)

	//return fmt.Sprintf("%s://%s@%s:%s",
	//	protocol,
	//	getUserAndPassword(username, password),
	//	hostname,
	//	string(port),
	//)
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