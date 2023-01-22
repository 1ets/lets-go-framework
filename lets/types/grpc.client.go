package types

import (
	"lets-go-framework/lets"
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
	GetClients() []IGrpcServiceClient
}

// Client information
type GrpcClient struct {
	Name    string
	Host    string
	Port    string
	Clients []IGrpcServiceClient
}

// Get Name
func (gc *GrpcClient) GetName() string {
	if gc.Name == "" {
		lets.LogW("Configs: CLIENT_GRPC_NAME is not set, using default configuration.")

		return CLIENT_GRPC_NAME
	}

	return gc.Name
}

// Get Host
func (gc *GrpcClient) GetHost() string {
	if gc.Host == "" {
		lets.LogW("Configs: CLIENT_GRPC_HOST is not set, using default configuration.")

		return CLIENT_GRPC_HOST
	}

	return gc.Host
}

// Get Port
func (gc *GrpcClient) GetPort() string {
	if gc.Port == "" {
		lets.LogW("Configs: CLIENT_GRPC_PORT is not set, using default configuration.")

		return CLIENT_GRPC_PORT
	}

	return gc.Port
}

// Get Clients
func (gc *GrpcClient) GetClients() []IGrpcServiceClient {
	return gc.Clients
}
