package clients

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
)

// Define RabbitMQ client, it will used by controller.
var RabbitSagaExampleCallback = &rabbitSagaExampleCallback{}

// RabbitMQ client definition.
type rabbitSagaExampleCallback struct {
	c types.IFrameworkRabbitMQPublisher
}

// Its implementation of types.IGrpcServiceClient.
func (g *rabbitSagaExampleCallback) SetConnection(c types.IFrameworkRabbitMQPublisher) {
	g.c = c
}

// Reply callback for transaction service provider.
func (g *rabbitSagaExampleCallback) TransactionCallback(r types.IEvent, data *data.ResponseTransfer) (err error) {
	var eventName = "transfer-callback"

	// Publish callback where its dont have replyTo.
	if err = g.c.Publish(&types.Event{
		Debug:         true,
		Name:          eventName,
		Data:          data,
		Exchange:      r.GetReplyTo().Exchange,
		RoutingKey:    r.GetReplyTo().RoutingKey,
		CorrelationId: r.GetCorrelationId(),
	}); err != nil {
		lets.LogE(err.Error())
		return
	}

	return
}

// Reply callback for balance service provider.
func (g *rabbitSagaExampleCallback) BalanceCallback(r types.IEvent, data *data.ResponseBalance) (err error) {
	var eventName = "balance-callback"

	// Publish callback where its dont have replyTo.
	if err = g.c.Publish(&types.Event{
		Debug:         true,
		Name:          eventName,
		Data:          data,
		Exchange:      r.GetReplyTo().Exchange,
		RoutingKey:    r.GetReplyTo().RoutingKey,
		CorrelationId: r.GetCorrelationId(),
	}); err != nil {
		lets.LogE(err.Error())
		return
	}

	return
}

// Reply callback for balance service provider.
func (g *rabbitSagaExampleCallback) NotificationCallback(r types.IEvent, data *data.ResponseNotification) (err error) {
	var eventName = "notification-callback"

	// Publish callback where its dont have replyTo.
	if err = g.c.Publish(&types.Event{
		Debug:         true,
		Name:          eventName,
		Data:          data,
		Exchange:      r.GetReplyTo().Exchange,
		RoutingKey:    r.GetReplyTo().RoutingKey,
		CorrelationId: r.GetCorrelationId(),
	}); err != nil {
		lets.LogE(err.Error())
		return
	}

	return
}
