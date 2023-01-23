package types

import (
	"lets-go-framework/lets"

	"google.golang.org/grpc"
)

// Default grpc server configuration
const (
	SERVER_GRPC_PORT = "5100"
)

// Interface for gRPC
type IGrpcServer interface {
	GetPort() string
	GetRouter() func(*grpc.Server)
}

// Server information
type GrpcServer struct {
	Port   string
	Router func(*grpc.Server)
}

// Get Port
func (g *GrpcServer) GetPort() string {
	if g.Port == "" {
		lets.LogW("Config: SERVER_GRPC_PORT is not set, using default configuration.")

		return SERVER_GRPC_PORT
	}

	return g.Port
}

// Get Router
func (g *GrpcServer) GetRouter() func(*grpc.Server) {
	return g.Router
}
