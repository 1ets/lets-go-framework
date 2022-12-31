package boot

import (
	"lets-go-framework/adapters"
	"lets-go-framework/lets/drivers"
)

func LoadGrpcFramework() {
	grpcDriver := drivers.NewGrpc()
	adapters.GrpcService(grpcDriver)
	grpcDriver.Connect()
	adapters.GrpcClient(grpcDriver)
}
