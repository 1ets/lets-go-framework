package framework

import (
	"fmt"
	"lets-go-framework/adapters"
	"lets-go-framework/lets/drivers"
)

func Grpc() {
	fmt.Println("LoadGrpcFramework()")

	grpcDriver := drivers.NewGrpc()
	adapters.GrpcService(grpcDriver)
	grpcDriver.Connect()
	adapters.GrpcClient(grpcDriver)
}
