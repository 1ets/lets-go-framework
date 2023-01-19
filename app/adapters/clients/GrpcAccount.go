package clients

import (
	"lets-go-framework/app/adapters/protobuf"

	"google.golang.org/grpc"
)

var GrpcAccount protobuf.ApiAccountClient

func SetGrpcAccount(gc *grpc.ClientConn) {
	GrpcAccount = protobuf.NewApiAccountClient(gc)
}
