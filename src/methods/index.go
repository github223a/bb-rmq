package methods

import (
	"../entities"
)

var List = map[string] *entities.Method {
	"friendship": Friendship,
	"infrastructure": Infrastructure,
}
