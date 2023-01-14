package configs

import "lets-go-framework/lets/framework"

var OnInits = []func(){
	InitializeRabbitMQ,
}

var OnMains = []func(){
	framework.Http,
	framework.Grpc,
	framework.RabbitMQ,
}
