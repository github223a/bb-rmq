package entities

var Emitter *EventEmitter

type EventEmitter struct {
	Channels map[string] chan interface{} `json:"channels"`
}

func CreateEmitter() *EventEmitter {
	var channels = make(map[string] chan interface{})
	return &EventEmitter{
		Channels: channels,
	}
}
