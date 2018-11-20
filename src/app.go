package src

import (
	methods "./internal-methods"
	"fmt"
	"github.com/streadway/amqp"
)

func processingInternalMethod(channel *amqp.Channel, parsedMessage map[string]interface{}) {
	validateRequest(parsedMessage)
	checkNamespace(parsedMessage)
	checkInternalMethod(parsedMessage)
	checkToken(parsedMessage)

	runMethod(channel, parsedMessage)
}

func validateRequest(parsedMessage map[string]interface{}) {
}
func checkNamespace(parsedMessage map[string]interface{}) {
}
func checkInternalMethod(parsedMessage map[string]interface{}) {
}
func checkToken(parsedMessage map[string]interface{}) {
}
func runMethod(channel *amqp.Channel, parsedMessage map[string]interface{}) {
	name := parsedMessage["method"].(string)

	if methods.List[name] == nil {
		fmt.Println("no method")
		return // need send error to client
	}
	methods.List[name](channel)
}

func sendResponseToClient() {

}