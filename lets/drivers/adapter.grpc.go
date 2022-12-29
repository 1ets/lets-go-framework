package drivers

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcEngine struct {
	Connections []*GrpcConnection
}

type GrpcConnection struct {
	Name       string
	Host       string
	Port       string
	Connection *grpc.ClientConn
	Error      error
}

func NewGrpc() *GrpcEngine {
	return &GrpcEngine{}
}

func (gc *GrpcConnection) connect() {
	server := fmt.Sprintf("%s:%s", gc.Host, gc.Port)

	gc.Connection, gc.Error = grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("Connect to ", server)
	if gc.Error != nil {
		fmt.Println("Connect to ", gc.Error)
		gc.Error = nil
		return
	}
}

func (pool *GrpcEngine) AddService(name string, host string, port string) {
	pool.Connections = append(pool.Connections, &GrpcConnection{
		Name: name,
		Host: host,
		Port: port,
	})

}
func (pool *GrpcEngine) Connect() {
	for _, gc := range pool.Connections {
		gc.connect()
	}
}

func (pool *GrpcEngine) GetService(name string) grpc.ClientConnInterface {
	for _, gb := range pool.Connections {
		if gb.Name == name {
			return gb.Connection
		}
	}

	return nil
}
