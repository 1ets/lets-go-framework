package services

import (
	"lets-go-framework/adapters/protobuf"
	"lets-go-framework/adapters/servers"

	"google.golang.org/grpc"
)

func RouteGrpcService(gs *grpc.Server) {
	protobuf.RegisterApiAccountServer(gs, &servers.ApiAccountServer{})
}
