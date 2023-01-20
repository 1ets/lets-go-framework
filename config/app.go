package config

import (
	"lets-go-framework/lets/frameworks"
	"lets-go-framework/lets/types"
	"lets-go-framework/services"
	"os"
)

// Load configuration into the kernel
func App() {
	frameworks.HttpConfig = &types.HttpServer{
		Port:       os.Getenv("SERVE_HTTP_PORT"),
		Middleware: services.HttpMiddleware,
		Router:     services.HttpRouter,
	}

	// frameworks.GrpcServerConfig = &types.GrpcServer{
	// 	Host:   os.Getenv("SERVE_GRPC_PORT"),
	// 	Port:   os.Getenv("SERVE_GRPC_PORT"),
	// 	Router: services.RouteGrpcService,
	// }

	// frameworks.GrpcClientConfig = []types.GrpcClient{
	// 	{
	// 		Name:   "account",
	// 		Host:   "127.0.0.1",
	// 		Port:   "5100",
	// 		Client: clients.SetGrpcAccount,
	// 	},
	// }

	// frameworks.RabbitMQDsn = &types.RabbitMQDsn{
	// 	Host:        os.Getenv("RQ_HOST"),
	// 	Port:        os.Getenv("RQ_PORT"),
	// 	Username:    os.Getenv("RQ_USERNAME"),
	// 	Password:    os.Getenv("RQ_PASSWORD"),
	// 	VirtualHost: os.Getenv("RQ_VHOST"),
	// }

	// frameworks.RabbitMQConsumer = &types.RabbitMQConsumer{
	// 	Name:         os.Getenv("RQ_NAME"),
	// 	Exchange:     os.Getenv("RQ_EXCHANGE"),
	// 	ExchangeType: amqp091.ExchangeDirect,
	// 	RoutingKey:   os.Getenv("RQ_ROUTING_KEY"),
	// 	Queue:        os.Getenv("RQ_QUEUE"),
	// }
}
