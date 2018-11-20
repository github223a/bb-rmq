package internal_methods

var List = map[string] func() {
	"friendship": handshake,
	"infrastructure": infrastructure,
}
