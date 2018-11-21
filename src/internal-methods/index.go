package internal_methods

import "../templates"

var List = map[string] func(templates.Request) {
	"friendship": handshake,
	"infrastructure": infrastructure,
}
