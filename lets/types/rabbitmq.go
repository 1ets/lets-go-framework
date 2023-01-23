package types

import "lets-go-framework/lets"

type IRabbitMQConfig interface {
	GetServers() []IRabbitMQServer
}

type RabbitMQConfig struct {
	Servers []IRabbitMQServer
}

func (r *RabbitMQConfig) GetServers() []IRabbitMQServer {
	return r.Servers
}

type IRabbitMQServer interface {
	GetHost() string
	GetPort() string
	GetUsername() string
	GetPassword() string
	GetVHost() string
	GetConsumers() []IRabbitMQConsumer
	GetPublishers() []IRabbitMQPublisher
}

// Default configuration
const (
	RABBIT_HOST     = "localhost"
	RABBIT_PORT     = "5672"
	RABBIT_USERNAME = "guest"
	RABBIT_PASSWORD = "guest"
	RABBIT_VHOST    = "/"
)

// Target host information.
type RabbitMQServer struct {
	Host        string
	Port        string
	Username    string
	Password    string
	VirtualHost string
	Consumers   []IRabbitMQConsumer
	Publishers  []IRabbitMQPublisher
}

// Get Host.
func (r *RabbitMQServer) GetHost() string {
	if r.Host == "" {
		lets.LogW("Config: RABBIT_HOST is not set, using default configuration.")

		return RABBIT_HOST
	}

	return r.Host
}

// Get Port.
func (r *RabbitMQServer) GetPort() string {
	if r.Port == "" {
		lets.LogW("Config: RABBIT_PORT is not set, using default configuration.")

		return RABBIT_PORT
	}

	return r.Port
}

// Get Username.
func (r *RabbitMQServer) GetUsername() string {
	if r.Username == "" {
		lets.LogW("Config: RABBIT_USERNAME is not set, using default configuration.")

		return RABBIT_USERNAME
	}

	return r.Username
}

// Get Password.
func (r *RabbitMQServer) GetPassword() string {
	if r.Password == "" {
		lets.LogW("Config: RABBIT_PASSWORD is not set, using default configuration.")

		return RABBIT_PASSWORD
	}

	return r.Password
}

// Get Virtual Host.
func (r *RabbitMQServer) GetVHost() string {
	if r.VirtualHost == "" {
		lets.LogW("Config: RABBIT_VHOST is not set, using default configuration.")

		return RABBIT_VHOST
	}

	return r.VirtualHost
}

// Get Consumers.
func (r *RabbitMQServer) GetConsumers() []IRabbitMQConsumer {
	return r.Consumers
}

// Get Publisher.
func (r *RabbitMQServer) GetPublishers() []IRabbitMQPublisher {
	return r.Publishers
}
