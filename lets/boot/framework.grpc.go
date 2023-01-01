package boot

import (
	"fmt"
	"lets-go-framework/adapters"
	"lets-go-framework/lets/drivers"
)

func LoadGrpcFramework() {
	fmt.Println("LoadGrpcFramework()")

	grpcDriver := drivers.NewGrpc()
	adapters.GrpcService(grpcDriver)
	grpcDriver.Connect()
	adapters.GrpcClient(grpcDriver)
}
