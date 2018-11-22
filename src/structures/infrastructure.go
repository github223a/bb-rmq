package structures

type InfrastructureData struct {
	RedisPrefix string `json:"redisPrefix"`
	RedisPrefixSession string `json:"redisPrefixSession"`
	RedisPrefixSessionList string `json:"redisPrefixSessionList"`
	TokenAlg string `json:"tokenAlg"`
	TokenKey string `json:"tokenKey"`
	SessionLifetime float64 `json:"sessionLifetime"`
	Expectation float64 `json:"expectation"`
	Shardings map[string] interface {} `json:"shardings"`
	Infrastructure map[string] InfrastructureServiceMethods `json:"infrastructure"`
	//Infrastructure map[string] interface{} `json:"infrastructure"`
}

type InfrastructureServiceMethods struct {
	Methods map[string] MethodSettings `json:"methods"`
}