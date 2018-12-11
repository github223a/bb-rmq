package main

type ConfigStrucuture struct {
	Namespace string `json:"namespace,omitempty"`
	// UseCache      bool     `json:"useCache"`
	// UseIsInternal bool     `json:"useIsInternal"`
	// Location      Location `json:"location"`
	// Log           core.Log      `json:"log"`
	// Redis         core.Redis    `json:"redis"`
	// RabbitMQ      core.RabbitMQ `json:"rabbitMQ"`
	// PingTimeout int32 `json:"pingTimeout"`
	// core.CoreConfig
}

type Location struct {
	Ws   Ws   `json:"ws"`
	Rest Rest `json:"rest"`
}

type Ws struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Path string `json:"path"`
}

type Rest struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Path string `json:"path"`
}
