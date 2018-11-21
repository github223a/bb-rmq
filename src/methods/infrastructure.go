package methods

import (
	"../constants"
	"../entities"
	"../templates"
	"fmt"
)

func runInfrastructure(request templates.Request) {
	constants.INFRASTRUCTURE = request.Params
	fmt.Println("infrastr ==== ", constants.INFRASTRUCTURE)
}

var infrastructureSettings = entities.MethodSettings{
	IsInternal: true,
	Auth: false,
	Cache: nil,
	Middlewares: entities.Middlewares {
		Before: []string{},
		After: []string{},
	},
}

var Infrastructure = entities.NewMethodEntity(runInfrastructure, infrastructureSettings)