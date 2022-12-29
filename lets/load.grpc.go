package lets

import (
	"lets-go-framework/adapters"
	"lets-go-framework/lets/drivers"
)

func loadGrpcFramework() {
	grpcDriver := drivers.NewGrpc()
	adapters.GrpcServiceConnection(grpcDriver)
	grpcDriver.Connect()
	adapters.GrpcServiceClient(grpcDriver)
}
