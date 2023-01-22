package clients

import (
	"context"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/adapters/protobuf"
	"lets-go-framework/lets"

	"google.golang.org/grpc"
)

// Define gRPC client, it will used by controller.
var GrpcExample = &grpcExample{}

// gRPC client definition.
type grpcExample struct {
	c protobuf.ExampleServiceClient
}

// Its implementation of types.IGrpcServiceClient.
func (g *grpcExample) SetConnection(c *grpc.ClientConn) {
	g.c = protobuf.NewExampleServiceClient(c)
}

// Send insert user data to gRPC server.
func (g *grpcExample) Greeting(push *data.RequestExample) (callback data.ResponseExample, err error) {
	var request protobuf.RequestExample
	lets.Bind(push, &request)

	var response = &protobuf.ResponseExample{}
	response, err = g.c.Example(context.Background(), &request)
	lets.Bind(response, &callback)

	callback.Greeting = response.GetData().GetGreeting()

	return
}
