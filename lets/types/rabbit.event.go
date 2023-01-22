package types

import (
	"encoding/json"
	"fmt"
)

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
