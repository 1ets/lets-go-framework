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
type IRabbitDsn interface {
	GetHost() string
	GetPort() string
	GetUsername() string
	GetPassword() string
	GetVirtualHost() string
}

// Target host information.
type RabbitDsn struct {
	Host, Port, Username, Password, VirtualHost string
}

// Get Host.
func (rtm *RabbitDsn) GetHost() string {
	if rtm.Host == "" {
		golog.Warn("Configs RabbitMQ: RQ_HOST is not set in .env file, using default configuration.")

		return RQ_HOST
	}

	return rtm.Host
}

// Get Port.
func (rtm *RabbitDsn) GetPort() string {
	if rtm.Port == "" {
		golog.Warn("Configs RabbitMQ: RQ_PORT is not set in .env file, using default configuration.")

		return RQ_PORT
	}

	return rtm.Port
}

// Get Username.
func (rtm *RabbitDsn) GetUsername() string {
	if rtm.Username == "" {
		golog.Warn("Configs RabbitMQ: RQ_USERNAME is not set in .env file, using default configuration.")

		return RQ_USERNAME
	}

	return rtm.Username
}

// Get Password.
func (rtm *RabbitDsn) GetPassword() string {
	if rtm.Password == "" {
		golog.Warn("Configs RabbitMQ: RQ_PASSWORD is not set in .env file, using default configuration.")

		return RQ_PASSWORD
	}

	return rtm.Password
}

// Get VirtualHost.
func (rtm *RabbitDsn) GetVirtualHost() string {
	if rtm.VirtualHost == "" {
		golog.Warn("Configs RabbitMQ: RQ_VHOST is not set in .env file, using default configuration.")

		return RQ_VHOST
	}

	return rtm.VirtualHost
}

// func NewRabbitDsn(host, port, username, password, virtualHost string) IRabbitDsn {
// 	return &RabbitDsn{
// 		Host:        host,
// 		Port:        port,
// 		Username:    username,
// 		Password:    password,
// 		VirtualHost: virtualHost,
// 	}
// }
