package types

// Engine for controller
type Engine interface {
	Event(string, func(*Event))
	Call(string, *Event)
}
