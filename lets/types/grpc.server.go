package types

import (
	"github.com/kataras/golog"
	"google.golang.org/grpc"
)

// Default grpc server configuration
const (
	SERVE_GRPC_HOST = "0.0.0.0"
	SERVE_GRPC_PORT = "5100"
)

// Interface for grpc method
type IGrpcServer interface {
	GetHost() string
	GetPort() string
	GetRouter() func(*grpc.Server)
}

// Server information
type GrpcServer struct {
	Host   string
	Port   string
	Router func(*grpc.Server)
}

// Get Host
func (hs *GrpcServer) GetHost() string {
	if hs.Port == "" {
		golog.Warn("Configs Http: GRPC_HOST is not set, using default configuration.")

		return SERVE_GRPC_HOST
	}

	return hs.Host
}

// Get Port
func (hs *GrpcServer) GetPort() string {
	if hs.Port == "" {
		golog.Warn("Configs Http: GRPC_PORT is not set, using default configuration.")

		return SERVE_GRPC_PORT
	}

	return hs.Port
}

// Get Router
func (hs *GrpcServer) GetRouter() func(*grpc.Server) {
	return hs.Router
}
