package templates

type Config struct {
	Namespace string `json:"namespace"`
	UseCache bool `json:"useCache"`
	UseIsInternal string `json:"useIsInternal"`
	Redis Redis `json:"redis"`
	Location Location `json:"location"`
	RabbitMQ RabbitMQ `json:"rabbitMQ"`
}

type RabbitMQ struct {
	Connection Connection `json:"connection"`
	Channels map[string] interface{} `json:"channels"`
	TestChannelName string `json:"testChannelName"`
	TestChannelPingTimeout int `json:"testChannelPingTimeout"`
	InfrastructureBindingKey string `json:"infrastructureBindingKey"`
}

type Connection struct {
	Protocol string `json:"protocol,omitempty" default:"amqp"`
	Hostname string `json:"hostname,omitempty" default:"localhost"`
	Username string `json:"username,omitempty" default:"guest"`
	Password string `json:"password,omitempty" default:"guest"`
	Port int `json:"port,omitempty" default:"5672"`
}

type Ws struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Path string `json:"path"`
}

type Rest struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Path string `json:"path"`
}

type Location struct {
	Ws Ws `json:"ws"`
	Rest Rest `json:"rest"`
}

type Redis struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Password string `json:"prefix"`
}
