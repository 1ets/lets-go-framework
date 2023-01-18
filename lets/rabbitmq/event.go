package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime"
)

// Engine for controller
type Engine struct {
	handlers []Handler
}

func (me *Engine) Event(name string, controller func(*Event)) {
	fmt.Printf("[RabbitMQ] Event %s --> %v\n", name, runtime.FuncForPC(reflect.ValueOf(controller).Pointer()).Name())

	me.handlers = append(me.handlers, Handler{
		Name:       name,
		Controller: controller,
	})
}

func (me *Engine) Call(name string, event *Event) {
	for _, handler := range me.handlers {
		if handler.Name == name {
			handler.Controller(event)
			return
		}
	}

	log.Default().Println("Event not found: ", name)
}

type Handler struct {
	Name       string
	Controller func(*Event)
}

// type MessageContext interface {
// 	GetEventName() string
// 	GetData() []byte
// 	GetReplyTo() string
// 	GetCorrelationId() string
// 	GetExchange() string
// 	GetRoutingKey() string
// }

type Event struct {
	Name          string
	Exchange      string
	RoutingKey    string
	Data          interface{}
	ReplyTo       ReplyTo
	CorrelationId string
}

func (m *Event) GetEventName() string {
	return m.Name
}

func (m *Event) GetData() interface{} {
	return m.Data
}

func (m *Event) GetReplyTo() ReplyTo {
	return m.ReplyTo
}

func (m *Event) GetCorrelationId() string {
	return m.CorrelationId
}

func (m *Event) GetExchange() string {
	return m.Exchange
}

func (m *Event) GetRoutingKey() string {
	return m.RoutingKey
}

type ReplyTo struct {
	Exchange      string `json:"exchange"`
	RoutingKey    string `json:"routing_key"`
	CorrelationId string `json:"correlation_id"`
}

func (r *ReplyTo) GetJson() string {
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Marshal ERR: ", err.Error())
	}

	return string(data)
}
