package entities

import "../templates"

type MethodSettings struct {
	IsInternal bool `json:"isInternal"`
	Auth bool `json:"auth"`
	Cache *int `json:"cache"`
	Middlewares Middlewares `json:"middlewares"`
}

type Middlewares struct {
	Before []string `json:"before"`
	After []string `json:"after"`
}

type Method struct {
	Run func(request templates.Request) `json:"run"`
	Settings MethodSettings `json:"settings"`
}

func NewMethodEntity(run func(request templates.Request), settings MethodSettings) *Method {
	return &Method {
		Run: run,
		Settings: settings,
	}
}