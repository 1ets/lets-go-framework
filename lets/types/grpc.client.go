package types

import (
	"github.com/kataras/golog"
	"google.golang.org/grpc"
)

// Default grpc configuration
const (
	CLIENT_GRPC_NAME = "lets-service"
	CLIENT_GRPC_HOST = "127.0.0.1"
	CLIENT_GRPC_PORT = "5100"
)

// Interface for grpc method
type IGrpcClient interface {
	GetName() string
	GetHost() string
	GetPort() string
}

// Client information
type GrpcClient struct {
	Name   string
	Host   string
	Port   string
	Client func(*grpc.ClientConn)
}

// Get Name
func (gc *GrpcClient) GetName() string {
	if gc.Port == "" {
		golog.Warn("Configs Http: CLIENT_GRPC_NAME is not set, using default configuration.")

		return CLIENT_GRPC_NAME
	}

	return gc.Host
}

// Get Host
func (gc *GrpcClient) GetHost() string {
	if gc.Port == "" {
		golog.Warn("Configs Http: CLIENT_GRPC_HOST is not set, using default configuration.")

		return CLIENT_GRPC_HOST
	}

	return gc.Host
}

// Get Port
func (gc *GrpcClient) GetPort() string {
	if gc.Port == "" {
		golog.Warn("Configs Http: CLIENT_GRPC_PORT is not set, using default configuration.")

		return CLIENT_GRPC_PORT
	}

	return gc.Port
}
