package configs

import (
	"lets-go-framework/libraries/rabbitmq"
	"os"
)

// Initialize global variables
var RabbitMQCredentials map[string]rabbitmq.ICredentials = map[string]rabbitmq.ICredentials{}

// Callable function for initialize all libraries configuration.
func Initialize() {
	// Rabbit: Order configuration setups.
	RabbitMQCredentials["order"] = &rabbitmq.Credentials{
		Host:        os.Getenv("RQ_HOST"),
		Port:        os.Getenv("RQ_PORT"),
		Username:    os.Getenv("RQ_USERNAME_ORDER"),
		Password:    os.Getenv("RQ_PASSWORD_ORDER"),
		VirtualHost: os.Getenv("RQ_SAGA_VHOST"),
		Exchange:    os.Getenv("RQ_EXCHANGE_ORDER"),
		RoutingKey:  os.Getenv("RQ_ROUTINGKEY_BOOKING"),
	}
}
