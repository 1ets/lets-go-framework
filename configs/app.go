package configs

import (
	"lets-go-framework/adapters/clients"
	"lets-go-framework/lets/framework"
	"lets-go-framework/lets/types"
	"lets-go-framework/services"
	"os"
)

func AppConfigs() {
	framework.HttpConfig = &types.HttpServer{
		Port:       os.Getenv("SERVE_HTTP_PORT"),
		Middleware: services.MiddlewareHttpService,
		Router:     services.RouteHttpService,
	}

	framework.GrpcServerConfig = &types.GrpcServer{
		Host:   os.Getenv("SERVE_GRPC_PORT"),
		Port:   os.Getenv("SERVE_GRPC_PORT"),
		Router: services.RouteGrpcService,
	}

	framework.GrpcClientConfig = []types.GrpcClient{
		types.GrpcClient{
			Name:   "account",
			Host:   "127.0.0.1",
			Port:   "5100",
			Client: clients.SetGrpcAccount,
		},
	}

	// framework.GrpcClientConfig = &types.GrpcClient{
	// 	Host: os.Getenv("CLIENT_GRPC_PORT"),
	// 	Port: os.Getenv("CLIENT_GRPC_PORT"),
	// }

	// framework.RabbitMQDsn = &types.RabbitDsn{
	// 	Host:        os.Getenv("RQ_HOST"),
	// 	Port:        os.Getenv("RQ_PORT"),
	// 	Username:    os.Getenv("RQ_USERNAME"),
	// 	Password:    os.Getenv("RQ_PASSWORD"),
	// 	VirtualHost: os.Getenv("RQ_VHOST"),
	// }

	// framework.RabbitMConsumer = &types.RabbitConsumer{
	// 	Name:         os.Getenv("RQ_NAME"),
	// 	Exchange:     os.Getenv("RQ_EXCHANGE"),
	// 	ExchangeType: amqp091.ExchangeDirect,
	// 	RoutingKey:   os.Getenv("RQ_ROUTING_KEY"),
	// 	Queue:        os.Getenv("RQ_QUEUE"),
	// }

}
