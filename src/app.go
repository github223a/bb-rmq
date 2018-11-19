package src
//
//import (
//	"./templates"
//	"encoding/json"
//	"fmt"
//	"github.com/streadway/amqp"
//	"log"
//)
//
//func SendHandhake() {
//	request := templates.Handshake()
//	jsonData, err := json.Marshal(request)
//
//	requestMsg := string(jsonData)
//
//	fmt.Println("json = ", requestMsg)
//	err = channel.Publish(
//		"",      // exchange
//		request.Namespace,           // routing key
//		false,  // mandatory
//		false,  // immediate
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(requestMsg),
//		})
//	log.Printf(" [x] Sent %s", requestMsg)
//	FailOnError(err, "Failed to publish a message")
//}