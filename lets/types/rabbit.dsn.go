package types

import (
	"github.com/kataras/golog"
)

// Default configuration
const (
	RQ_USERNAME = "guest"
	RQ_PASSWORD = "guest"
	RQ_HOST     = "localhost"
	RQ_PORT     = "5672"
	RQ_VHOST    = "/"
)

// Interface for dsn accessable method
type IRabbitMQDsn interface {
	GetHost() string
	GetPort() string
	GetUsername() string
	GetPassword() string
	GetVirtualHost() string
}

// Target host information.
type RabbitMQDsn struct {
	Host, Port, Username, Password, VirtualHost string
}

// Get Host.
func (rtm *RabbitMQDsn) GetHost() string {
	if rtm.Host == "" {
		golog.Warn("Configs RabbitMQ: RQ_HOST is not set in .env file, using default configuration.")

		return RQ_HOST
	}

	return rtm.Host
}

// Get Port.
func (rtm *RabbitMQDsn) GetPort() string {
	if rtm.Port == "" {
		golog.Warn("Configs RabbitMQ: RQ_PORT is not set in .env file, using default configuration.")

		return RQ_PORT
	}

	return rtm.Port
}

// Get Username.
func (rtm *RabbitMQDsn) GetUsername() string {
	if rtm.Username == "" {
		golog.Warn("Configs RabbitMQ: RQ_USERNAME is not set in .env file, using default configuration.")

		return RQ_USERNAME
	}

	return rtm.Username
}

// Get Password.
func (rtm *RabbitMQDsn) GetPassword() string {
	if rtm.Password == "" {
		golog.Warn("Configs RabbitMQ: RQ_PASSWORD is not set in .env file, using default configuration.")

		return RQ_PASSWORD
	}

	return rtm.Password
}

// Get VirtualHost.
func (rtm *RabbitMQDsn) GetVirtualHost() string {
	if rtm.VirtualHost == "" {
		golog.Warn("Configs RabbitMQ: RQ_VHOST is not set in .env file, using default configuration.")

		return RQ_VHOST
	}

	return rtm.VirtualHost
}
