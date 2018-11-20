package src

import (
	methods "./internal-methods"
	"fmt"
)

func processingInternalMethod(parsedMessage map[string]interface{}) {
	validateRequest(parsedMessage)
	checkNamespace(parsedMessage)
	checkInternalMethod(parsedMessage)
	checkToken(parsedMessage)

	runMethod(parsedMessage)
}

func validateRequest(parsedMessage map[string]interface{}) {
}
func checkNamespace(parsedMessage map[string]interface{}) {
}
func checkInternalMethod(parsedMessage map[string]interface{}) {
}
func checkToken(parsedMessage map[string]interface{}) {
}
func runMethod(parsedMessage map[string]interface{}) {
	name := parsedMessage["method"]
	fmt.Println(name)
	methods.List[name]()
	//methods[parsedMessage["method"]]
}