package adapters

import (
	"lets-go-framework/adapters/data"

	"github.com/rabbitmq/amqp091-go"
)

var EventTransaction = eventTransaction{}

type eventTransaction struct {
	Connection *amqp091.Connection
}

func (et *eventTransaction) PublishTransfer(request data.EventTransfer) {
	et.Connection
}

func (*eventTransaction) ConsumeTransferEvent() {

}
