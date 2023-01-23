package clients

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"os"
	"time"
)

// Define RabbitMQ client, it will used by controller.
var RabbitExample = &rabbitExample{}

// RabbitMQ client definition.
type rabbitExample struct {
	c              types.IFrameworkRabbitMQPublisher
	waiterGreeting map[string]chan types.Event // Bucket for stateless or synchronus
}

// Its implementation of types.IGrpcServiceClient.
func (g *rabbitExample) SetConnection(c types.IFrameworkRabbitMQPublisher) {
	g.c = c

}

// Request greeting to RabbitMQ server.
func (g *rabbitExample) GreetingSync(request *data.RequestExample) (response data.ResponseExample, err error) {
	var eventName = "example-event"

	// Wait reply
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())

	// Save channel
	if g.waiterGreeting == nil {
		g.waiterGreeting = map[string]chan types.Event{}
	}
	g.waiterGreeting[correlationId] = make(chan types.Event)
	defer delete(g.waiterGreeting, correlationId)

	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE"),
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY"),
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY"),
		},
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	var callback = <-g.waiterGreeting[correlationId]

	lets.Bind(callback.GetData(), &response)

	return
}

// Response handling for GreetingSync().
func (g *rabbitExample) GreetingSyncCallback(r *types.Event) {
	lets.LogD(lets.ToJson(r))
	if g.waiterGreeting[r.GetCorrelationId()] == nil {
		return
	}

	g.waiterGreeting[r.GetCorrelationId()] <- *r
}

// Request greeting to RabbitMQ server.
func (g *rabbitExample) GreetingAsync(request *data.RequestExample) (callback data.ResponseExample, err error) {
	var eventName = "example-event"

	if err = g.c.Publish(&types.Event{
		Debug:      true,
		Name:       eventName,
		Data:       request,
		Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE"),
		RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY"),
		ReplyTo: &types.ReplyTo{
			Exchange:   os.Getenv("LISTEN_RABBIT_EXCHANGE"),
			RoutingKey: os.Getenv("LISTEN_RABBIT_ROUTING_KEY"),
		},
	}); err != nil {
		return
	}

	callback.Greeting = "A message was sent via message broker, check your terminal."

	return
}

// Reply greeting to RabbitMQ sender.
func (g *rabbitExample) GreetingCallback(correlationId string, request *data.ResponseExample) (callback data.ResponseExample, err error) {
	var eventName = "callback"

	if err = g.c.Publish(&types.Event{
		Debug:         true,
		Name:          eventName,
		Data:          request,
		Exchange:      os.Getenv("LISTEN_RABBIT_EXCHANGE"),
		RoutingKey:    os.Getenv("LISTEN_RABBIT_ROUTING_KEY"),
		CorrelationId: correlationId,
	}); err != nil {
		return
	}

	callback.Greeting = "A reply message was sent."

	return
}
