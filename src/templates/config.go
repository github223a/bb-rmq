package templates

type Connection struct {
	Protocol string `json:"protocol,omitempty" default:"amqp"`
	Hostname string `json:"hostname,omitempty" default:"localhost"`
	Username string `json:"username,omitempty" default:"guest"`
	Password string `json:"password,omitempty" default:"guest"`
	Port int `json:"port,omitempty" default:"5672"`
}

type Config struct {
	Connection Connection `json:"connection"`
	Channels map[string] interface{} `json:"channels"`
	TestChannelName string `json:"testChannelName"`
	TestChannelPingTimeout int `json:"testChannelPingTimeout"`
	InfrastructureBindingKey string `json:"infrastructureBindingKey"`
}
