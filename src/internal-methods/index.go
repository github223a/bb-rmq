package internal_methods

type fn func ()

var List = map[string] fn {
	"handshake": handshake,
	"infrastructure": infrastructure,
}
