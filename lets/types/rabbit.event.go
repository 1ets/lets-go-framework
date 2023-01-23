package types

import (
	"encoding/json"
	"fmt"
	"lets-go-framework/lets"
)

type IEvent interface {
	GetName() string
	GetData() interface{}
	GetReplyTo() *ReplyTo
	GetCorrelationId() string
	GetExchange() string
	GetRoutingKey() string
	GetBody() []byte
	GetDebug() bool
}

type Event struct {
	Name          string
	Exchange      string // Service exchange.
	RoutingKey    string // Service routing key.
	Data          interface{}
	ReplyTo       *ReplyTo
	CorrelationId string
	Debug         bool
}

func (m *Event) GetName() string {
	return m.Name
}

func (m *Event) GetData() interface{} {
	return m.Data
}

func (m *Event) GetReplyTo() *ReplyTo {
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

func (m *Event) GetDebug() bool {
	return m.Debug
}

func (m *Event) GetBody() []byte {
	body, err := json.Marshal(RabbitBody{Event: m.Name, Data: m.Data})
	if err != nil {
		lets.LogE("RabbitEvent: %s", err.Error())
		return nil
	}

	return body
}

type ReplyTo struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
}

func (r *ReplyTo) GetJson() string {
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Marshal ERR: ", err.Error())
	}

	return string(data)
}
