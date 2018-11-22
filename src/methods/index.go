package methods

import (
	"../structures"
)

var List = map[string] *structures.Method {
	"friendship": friendship,
	"infrastructure": infrastructure,
}

func NewMethodEntity(run func(request structures.Request), settings structures.MethodSettings) *structures.Method {
	return &structures.Method {
		Run: run,
		Settings: settings,
	}
}
