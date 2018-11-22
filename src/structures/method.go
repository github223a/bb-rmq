package structures

type Method struct {
	Run func(request Request) `json:"run"`
	Settings MethodSettings `json:"settings"`
}

type MethodSettings struct {
	IsInternal bool `json:"isInternal"`
	Auth bool `json:"auth"`
	Cache int `json:"cache"`
	Middlewares Middlewares `json:"middlewares"`
}

type Middlewares struct {
	Before []string `json:"before"`
	After []string `json:"after"`
}

