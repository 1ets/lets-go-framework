package framework

import (
	"fmt"
	"lets-go-framework/lets/types"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GrpcServerConfig types.IGrpcServer
var GrpcClientConfig []types.GrpcClient

// 	grpcDriver := drivers.NewGrpc()
// 	adapters.GgrpcService(grpcDriver)
// 	grpcDriver.Connect()
// 	adapters.GrpcClient(grpcDriver)
// }

type grpcService struct {
	Server string
	Engine *grpc.Server
	Router func(*grpc.Server)
}

func (rpc *grpcService) Init() {
	fmt.Println("grpcService.Init()")

	var opts []grpc.ServerOption

	rpc.Server = fmt.Sprintf("%s:%s", GrpcServerConfig.GetHost(), GrpcServerConfig.GetPort())
	rpc.Engine = grpc.NewServer(opts...)
	rpc.Router = GrpcServerConfig.GetRouter()
}

func (rpc *grpcService) Serve() {
	fmt.Println("grpcService.Serve()")

	listener, err := net.Listen("tcp", rpc.Server)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpc.Engine.Serve(listener)
}

type grpcClient struct {
	Server string
	Engine *grpc.ClientConn
}

func (rpc *grpcClient) Init() {
	fmt.Println("grpcService.Init()")

	for _, config := range GrpcClientConfig {
		dsn = fmt.Sprintf("%s:%s", config.GetHost(), config.GetPort())
		var err error

		conn, err = grpc.Dial(dsn, grpc.WithTransportCredentials(insecure.NewCredentials()))
		fmt.Println("Connect to ", dsn)
		if err != nil {
			fmt.Println("Error ", err.Error())
			return
		}

		config.Client(conn)
	}
}

func (rpc *grpcClient) Connect() {
	fmt.Println("grpcService.Serve()")

}

// Define rpcservice host and port
func Grpc() {
	fmt.Println("Grpc()")

	var rpcServer grpcService
	rpcServer.Init()
	rpcServer.Router(rpcServer.Engine)
	rpcServer.Serve()

	var rpcClient grpcClient
	rpcClient.Init()
	rpcClient.Connect()
}
