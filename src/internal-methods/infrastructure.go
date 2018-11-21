package internal_methods

import (
	"../constants"
	"../templates"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func infrastructure(request templates.Request) {
	err := mapstructure.Decode(request.Params, &constants.INFRASTRUCTURE)
	if err != nil {
		panic(err)
	}
	fmt.Println("infrastr ==== ", constants.INFRASTRUCTURE)
}