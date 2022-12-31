package drivers

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverGrpc struct {
	Connections []*Grpc
}

type Grpc struct {
	Name       string
	Host       string
	Port       string
	Connection *grpc.ClientConn
	Error      error
}

func NewGrpc() *DriverGrpc {
	return &DriverGrpc{}
}

func (gc *Grpc) connect() {
	server := fmt.Sprintf("%s:%s", gc.Host, gc.Port)

	gc.Connection, gc.Error = grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("Connect to ", server)
	if gc.Error != nil {
		fmt.Println("Connect to ", gc.Error)
		gc.Error = nil
		return
	}
}

func (pool *DriverGrpc) AddService(name string, host string, port string) {
	pool.Connections = append(pool.Connections, &Grpc{
		Name: name,
		Host: host,
		Port: port,
	})

}
func (pool *DriverGrpc) Connect() {
	for _, gc := range pool.Connections {
		gc.connect()
	}
}

func (pool *DriverGrpc) GetService(name string) grpc.ClientConnInterface {
	for _, gb := range pool.Connections {
		if gb.Name == name {
			return gb.Connection
		}
	}

	return nil
}
