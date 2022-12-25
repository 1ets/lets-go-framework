package rabbitmq

import (
	"github.com/kataras/golog"
)

// Default configuration
const (
	RQ_USERNAME   = "guest"
	RQ_PASSWORD   = "guest"
	RQ_HOST       = "localhost"
	RQ_PORT       = "5672"
	RQ_VHOST      = "/"
	RQ_EXCHANGE   = "default-exchange"
	RQ_ROUTINGKEY = "default-routingKey"
)

// Interface for credentials accessable method
type ICredentials interface {
	GetHost() string
	GetPort() string
	GetUsername() string
	GetPassword() string
	GetVirtualHost() string
	GetExchange() string
	GetRoutingKey() string
}

// Target host information.
type Credentials struct {
	Host, Port, Username, Password, VirtualHost, Exchange, RoutingKey string
}

// Get Host.
func (rtm *Credentials) GetHost() string {
	if rtm.Host == "" {
		golog.Warn("Configs RabbitMQ: RQ_HOST is not set in .env file, using default configuration.")

		return RQ_HOST
	}

	return rtm.Host
}

// Get Port.
func (rtm *Credentials) GetPort() string {
	if rtm.Port == "" {
		golog.Warn("Configs RabbitMQ: RQ_PORT is not set in .env file, using default configuration.")

		return RQ_PORT
	}

	return rtm.Port
}

// Get Username.
func (rtm *Credentials) GetUsername() string {
	if rtm.Username == "" {
		golog.Warn("Configs RabbitMQ: RQ_USERNAME is not set in .env file, using default configuration.")

		return RQ_USERNAME
	}

	return rtm.Username
}

// Get Password.
func (rtm *Credentials) GetPassword() string {
	if rtm.Password == "" {
		golog.Warn("Configs RabbitMQ: RQ_PASSWORD is not set in .env file, using default configuration.")

		return RQ_PASSWORD
	}

	return rtm.Password
}

// Get VirtualHost.
func (rtm *Credentials) GetVirtualHost() string {
	if rtm.VirtualHost == "" {
		golog.Warn("Configs RabbitMQ: RQ_VHOST is not set in .env file, using default configuration.")

		return RQ_VHOST
	}

	return rtm.VirtualHost
}

// Get ExchangeName.
func (rtm *Credentials) GetExchange() string {
	if rtm.Exchange == "" {
		golog.Warn("Configs RabbitMQ: RQ_EXCHANGE is not set in .env file, using default configuration.")

		return RQ_EXCHANGE
	}

	return rtm.Exchange
}

// Get QueueName.
func (rtm *Credentials) GetRoutingKey() string {
	if rtm.RoutingKey == "" {
		golog.Warn("Configs RabbitMQ: RQ_ROUTINGKEY is not set in .env file, using default configuration.")

		return RQ_ROUTINGKEY
	}

	return rtm.RoutingKey
}
