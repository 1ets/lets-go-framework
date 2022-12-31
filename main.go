package main

import (
	"lets-go-framework/configs"
	"lets-go-framework/lets"
	"lets-go-framework/lets/boot"
)

// Initialize all required vars, consts

var bootstrap = lets.Bootstrap{
	OnInits: []func(){
		boot.LoadEnv,
		configs.Initialize,
	},
	OnMains: []func(){
		boot.LoadHttpFramework,
		boot.LoadGrpcFramework,
	},
}

func init() {
	bootstrap.OnInit()
}

func main() {
	bootstrap.OnMain()
}
