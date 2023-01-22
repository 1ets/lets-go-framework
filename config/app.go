package config

import (
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/lets/frameworks"
	"lets-go-framework/lets/types"
	"lets-go-framework/services"
	"os"
)

// Load configuration into the kernel
func App() {
	// Setup HTTP server
	frameworks.HttpConfig = &types.HttpServer{
		Port:       os.Getenv("SERVER_HTTP_PORT"),
		Mode:       os.Getenv("SERVER_HTTP_MODE"),
		Middleware: services.HttpMiddleware,
		Router:     services.HttpRouter,
	}

	// Setup gRPC
	frameworks.GrpcConfig = &types.GrpcConfig{
		// Setup server
		Server: &types.GrpcServer{
			Port:   os.Getenv("SERVER_GRPC_PORT"),
			Router: services.GrpcRouter,
		},

		// Setup gRPC Client
		Clients: []types.IGrpcClient{
			&types.GrpcClient{
				Name: os.Getenv("CLIENT_GRPC_NAME_EXAMPLE"),
				Host: os.Getenv("CLIENT_GRPC_HOST_EXAMPLE"),
				Port: os.Getenv("CLIENT_GRPC_PORT_EXAMPLE"),
				Clients: []types.IGrpcServiceClient{
					clients.GrpcExample,
				},
			},
		},
	}


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
