package drivers

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

var MessageEngines = MessageEngine{}

type MessageEngine struct {
	events []MessageEvent
}

func (me *MessageEngine) Event(name string, handler MessageHandler) {
	fmt.Printf("[RabbitMQ] Event %s --> %v", name, runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())

	me.events = append(me.events, MessageEvent{
		Name:    name,
		Handler: handler,
	})
}

func (me *MessageEngine) Call(name string, message MessageContext) {
	for _, eventhandler := range me.events {
		if eventhandler.Name == name {
			eventhandler.Handler(message)
			return
		}
	}

	log.Default().Println("Event not found: ", name)
}

type MessageEvent struct {
	Name    string
	Handler MessageHandler
}

type MessageHandler func(MessageContext)

type MessageContext interface {
	GetEventName() string
	GetData() []byte
	GetReplyTo() string
	GetCorrelationId() string
	GetExchange() string
	GetRoutingKey() string
}
