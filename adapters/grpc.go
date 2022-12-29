package adapters

import (
	"lets-go-framework/adapters/protobuf"
	"lets-go-framework/lets/drivers"
)

func GrpcServiceConnection(pool *drivers.GrpcEngine) {
	pool.AddService("transaction", "127.0.0.1", "5100")
	pool.AddService("account", "127.0.0.1", "5100")
}

var GrpcTransaction protobuf.TransactionServiceClient
var GrpcAccount protobuf.AccountServiceClient

func GrpcServiceClient(pool *drivers.GrpcEngine) {
	GrpcTransaction = protobuf.NewTransactionServiceClient(pool.GetService("transaction"))
	GrpcAccount = protobuf.NewAccountServiceClient(pool.GetService("account"))
}
