package types

import (
	"lets-go-framework/lets"
)

// Default configuration
const (
	CALLER_RABBIT_NAME_EXAMPLE     = "default-name"
	CALLER_RABBIT_EXCHANGE_EXAMPLE = "default-exchange"
	CALLER_RABBIT_DEBUG_EXAMPLE    = false
)

// Interface for dsn accessable method
type IRabbitMQPublisher interface {
	GetName() string
	GetClients() []IRabbitMQServiceClient
}

// Target host information.
type RabbitMQPublisher struct {
	Name    string `json:"name"`
	Clients []IRabbitMQServiceClient
}

// Get ExchangeName.
func (r *RabbitMQPublisher) GetName() string {
	if r.Name == "" {
		lets.LogW("Configs: CALLER_RABBIT_NAME_EXAMPLE is not set, using default configuration.")

		return CALLER_RABBIT_NAME_EXAMPLE
	}

	return r.Name
}

// Get Clients.
func (r *RabbitMQPublisher) GetClients() []IRabbitMQServiceClient {
	return r.Clients
}

type IRabbitMQServiceClient interface {
	SetConnection(IFrameworkRabbitMQPublisher)
}

type IFrameworkRabbitMQPublisher interface {
	Publish(IEvent) error
}
