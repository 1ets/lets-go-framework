package types

import "google.golang.org/grpc"

// GRPC configuration interface
type IGrpcConfig interface {
	GetServer() IGrpcServer
	GetClients() []IGrpcClient
}

// GRPC configuration struct
type GrpcConfig struct {
	Server  IGrpcServer
	Clients []IGrpcClient
}

// Get gRPC server configuration
func (g *GrpcConfig) GetServer() IGrpcServer {
	return g.Server
}

// Get gRPC client configuration
func (g *GrpcConfig) GetClients() []IGrpcClient {
	return g.Clients
}

type IGrpcServiceClient interface {
	SetConnection(*grpc.ClientConn)
}
