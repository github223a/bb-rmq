package rmq

import (
	"../templates"
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
	port, portOk := connection["port"].(int)

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

func Init() {
	url := getRabbitUrl()
	fmt.Println(url)
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	channels := rabbitMQ["channels"].(map[string] interface{})

	for name, _ := range channels {
		fmt.Println(name)
	}

	args := make(amqp.Table)
	args["x-message-ttl"] = int32(30000)

	q, err := ch.QueueDeclare(
		"internal", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		args,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	request := templates.Handshake()
	jsonData, err := json.Marshal(request)
	requestMsg := string(jsonData)

	fmt.Println("json = ", requestMsg)

	err = ch.Publish(
		"",      // exchange
		q.Name,           // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(requestMsg),
		})
	log.Printf(" [x] Sent %s", requestMsg)
	failOnError(err, "Failed to publish a message")
}