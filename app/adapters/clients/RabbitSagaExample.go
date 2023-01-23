package clients

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"os"
	"time"
)

// This section provides an example of simulating a saga transaction using its own service,
// maybe in reality the services will stand alone. Normally the service consists of
// transaction, balance and notification providers.

// Define RabbitMQ client, it will used by controller.
var RabbitSagaExample = &rabbitSagaExample{}

// RabbitMQ client definition.
type rabbitSagaExample struct {
	c                 types.IFrameworkRabbitMQPublisher
	wTransfer         map[string]chan types.Event // Bucket for stateless or synchronus, Transfer
	wTransferRollback map[string]chan types.Event // Bucket for stateless or synchronus, Transfer Rollback
	wBalance          map[string]chan types.Event // Bucket for stateless or synchronus, Balance
	wBalanceRollback  map[string]chan types.Event // Bucket for stateless or synchronus, Balance Rollback
	wNotification     map[string]chan types.Event // Bucket for stateless or synchronus, Notification
}

// Its implementation of types.IGrpcServiceClient.
func (g *rabbitSagaExample) SetConnection(c types.IFrameworkRabbitMQPublisher) {
	g.c = c
}

// Request transfer to transfer service provider.
func (g *rabbitSagaExample) Transfer(request *data.RequestTransfer) (response data.ResponseTransfer, err error) {
	var eventName = "transfer-request"

	if g.wTransfer == nil {
		g.wTransfer = map[string]chan types.Event{}
	}

	// The id for callback.
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	g.wTransfer[correlationId] = make(chan types.Event)
	defer delete(g.wTransfer, correlationId)

	// Start publishing data
	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),    // Exchange where transfer service provider binds.
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"), // RoutingKey for transfer service provider.

		// Tell transfer service provider where they must to be reply.
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"),
		},
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	// Wait and bind callback and return response to saga.
	var callback = <-g.wTransfer[correlationId]
	lets.Bind(callback.GetData(), &response)

	return
}

// Response handler for callback from transfer service provider.
func (g *rabbitSagaExample) TransferCallback(r *types.Event) {
	if g.wTransfer[r.GetCorrelationId()] == nil {
		return
	}

	// Send back event to waiter.
	g.wTransfer[r.GetCorrelationId()] <- *r
}

// Request transfer to transfer service provider.
func (g *rabbitSagaExample) TransferRollback(request *data.RequestTransferRollback) (response data.ResponseTransferRollback, err error) {
	var eventName = "transfer-rollback"

	if g.wTransferRollback == nil {
		g.wTransferRollback = map[string]chan types.Event{}
	}

	// The id for callback.
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	g.wTransferRollback[correlationId] = make(chan types.Event)
	defer delete(g.wTransferRollback, correlationId)

	// Start publishing data
	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),    // Exchange where transfer service provider binds.
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"), // RoutingKey for transfer service provider.

		// Tell transfer service provider where they must to be reply.
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"),
		},
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	// Wait and bind callback and return response to saga.
	var callback = <-g.wTransferRollback[correlationId]
	lets.Bind(callback.GetData(), &response)

	return
}

// Response handler for callback from transfer service provider.
func (g *rabbitSagaExample) TransferRollbackCallback(r *types.Event) {
	if g.wTransferRollback[r.GetCorrelationId()] == nil {
		return
	}

	// Send back event to waiter.
	g.wTransferRollback[r.GetCorrelationId()] <- *r
}

// Request transfer to balance service provider.
func (g *rabbitSagaExample) Balance(request *data.RequestBalance) (response data.ResponseBalance, err error) {
	var eventName = "balance-request"

	if g.wBalance == nil {
		g.wBalance = map[string]chan types.Event{}
	}

	// The id for callback.
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	g.wBalance[correlationId] = make(chan types.Event)
	defer delete(g.wBalance, correlationId)

	// Start publishing data
	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),    // Exchange where transfer service provider binds.
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"), // RoutingKey for transfer service provider.

		// Tell transfer service provider where they must to be reply.
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"),
		},
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	// Wait and bind callback and return response to saga.
	var callback = <-g.wBalance[correlationId]
	lets.Bind(callback.GetData(), &response)

	return
}

// Response handler for callback from balance service provider.
func (g *rabbitSagaExample) BalanceCallback(r *types.Event) {
	if g.wBalance[r.GetCorrelationId()] == nil {
		return
	}

	// Send back event to waiter.
	g.wBalance[r.GetCorrelationId()] <- *r
}

// Request transfer to transfer service provider.
func (g *rabbitSagaExample) BalanceRollback(request *data.RequestBalance) (response data.ResponseBalance, err error) {
	var eventName = "balance-rollback"

	if g.wBalanceRollback == nil {
		g.wBalanceRollback = map[string]chan types.Event{}
	}

	// The id for callback.
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	g.wBalanceRollback[correlationId] = make(chan types.Event)
	defer delete(g.wBalanceRollback, correlationId)

	// Start publishing data
	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),    // Exchange where transfer service provider binds.
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"), // RoutingKey for transfer service provider.

		// Tell transfer service provider where they must to be reply.
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"),
		},
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	// Wait and bind callback and return response to saga.
	var callback = <-g.wBalanceRollback[correlationId]
	lets.Bind(callback.GetData(), &response)

	return
}

// Response handler for callback from transfer service provider.
func (g *rabbitSagaExample) BalanceRollbackCallback(r *types.Event) {
	if g.wBalanceRollback[r.GetCorrelationId()] == nil {
		return
	}

	// Send back event to waiter.
	g.wBalanceRollback[r.GetCorrelationId()] <- *r
}

// Request notification to notification service provider.
func (g *rabbitSagaExample) Notification(request *data.RequestNotification) (response data.ResponseNotification, err error) {
	var eventName = "notification-request"

	if g.wNotification == nil {
		g.wNotification = map[string]chan types.Event{}
	}

	// The id for callback.
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	g.wNotification[correlationId] = make(chan types.Event)
	defer delete(g.wNotification, correlationId)

	// Start publishing data
	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),    // Exchange where transfer service provider binds.
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"), // RoutingKey for transfer service provider.

		// Tell transfer service provider where they must to be reply.
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE_SAGA"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY_SAGA"),
		},
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	// Wait and bind callback and return response to saga.
	var callback = <-g.wNotification[correlationId]
	lets.Bind(callback.GetData(), &response)

	return
}

// Request notification to notification service provider.
func (g *rabbitSagaExample) NotificationCallback(r *types.Event) {
	if g.wNotification[r.GetCorrelationId()] == nil {
		return
	}

	// Send back event to waiter.
	g.wNotification[r.GetCorrelationId()] <- *r
}
