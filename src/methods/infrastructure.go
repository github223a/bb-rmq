package methods

import (
	"../constants"
	"../structures"
	"fmt"
	"log"
)

var infrastructure = NewMethodEntity(runInfrastructure, infrastructureMethodSettings)

func runInfrastructure(request structures.Request) {
	fmt.Printf("alallalaa %+v\n", request)

	constants.InfrastructureData = structures.InfrastructureData {
		RedisPrefix: request.Params["redisPrefix"].(string),
		RedisPrefixSession: request.Params["redisPrefixSession"].(string),
		RedisPrefixSessionList: request.Params["redisPrefixSessionList"].(string),
		TokenAlg: request.Params["tokenAlg"].(string),
		TokenKey: request.Params["tokenKey"].(string),
		SessionLifetime: request.Params["sessionLifetime"].(float64),
		Expectation: request.Params["expectation"].(float64),
		Shardings: request.Params["shardings"].(map[string] interface{}),
		Infrastructure: request.Params["infrastructure"].(map[string] structures.InfrastructureServiceMethods),
	}
	log.Printf("%sInfrastructure updated.", constants.HEADER_RMQ_MESSAGE)
}

var infrastructureMethodSettings = structures.MethodSettings {
	IsInternal: true,
	Auth: false,
	Cache: 0,
	Middlewares: structures.Middlewares {
		Before: []string{},
		After: []string{},
	},
}
