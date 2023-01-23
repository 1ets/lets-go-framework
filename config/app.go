package config

import (
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/lets/frameworks"
	"lets-go-framework/lets/types"
	"lets-go-framework/services"
	"os"

	"github.com/rabbitmq/amqp091-go"
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

	// Setup RabbitMQ
	frameworks.RabbitMQConfig = &types.RabbitMQConfig{

		// Possible to use multiple rabbit server.
		Servers: []types.IRabbitMQServer{

			// Setup RabbitMQ Server.
			&types.RabbitMQServer{
				Host:        os.Getenv("RABBIT_HOST"),
				Port:        os.Getenv("RABBIT_PORT"),
				Username:    os.Getenv("RABBIT_USERNAME"),
				Password:    os.Getenv("RABBIT_PASSWORD"),
				VirtualHost: os.Getenv("RABBIT_VHOST"),

				// Possible to create multiple consumer for multiple purpose.
				Consumers: []types.IRabbitMQConsumer{

					// Setup Consumers
					&types.RabbitMQConsumer{
						Name:         os.Getenv("LISTEN_RABBIT_NAME"),
						Exchange:     os.Getenv("LISTEN_RABBIT_EXCHANGE"),
						ExchangeType: amqp091.ExchangeDirect,
						RoutingKey:   os.Getenv("LISTEN_RABBIT_ROUTING_KEY"),
						Queue:        os.Getenv("LISTEN_RABBIT_QUEUE"),
						Debug:        os.Getenv("LISTEN_RABBIT_DEBUG"),
						Listener:     services.RabbitMQRouter,
					},
				},
				Publishers: []types.IRabbitMQPublisher{
					// Setup Publisher
					&types.RabbitMQPublisher{
						Name: os.Getenv("LISTEN_RABBIT_NAME"),
						Clients: []types.IRabbitMQServiceClient{
							clients.RabbitExample,
						},
					},
				},
			},
		},
	}
}
