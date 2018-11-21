package methods

import (
	"../constants"
	"../templates"
	"log"
)

func runInfrastructure(request templates.Request) {
	constants.INFRASTRUCTURE = request.Params
	log.Printf("%sInfrastructure updated.", constants.HEADER_RMQ_MESSAGE)
}

var infrastructureSettings = MethodSettings{
	IsInternal: true,
	Auth: false,
	Cache: 0,
	Middlewares: Middlewares {
		Before: []string{},
		After: []string{},
	},
}

var Infrastructure = NewMethodEntity(runInfrastructure, infrastructureSettings)