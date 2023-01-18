package boot

import (
	"fmt"
	"lets-go-framework/lets/framework"
	"lets-go-framework/lets/loader"
	"reflect"
	"runtime"
)

var Initializer = []func(){
	loader.Environment,
}

var Servers = []func(){
	framework.Http,
	framework.Grpc,
	framework.RabbitMQ,
}

func OnInit() {
	fmt.Println("Initialization")
	for i, initializer := range Initializer {
		fmt.Printf("%v. Initializing %s\n", i, runtime.FuncForPC(reflect.ValueOf(initializer).Pointer()).Name())
		initializer()
	}
}

func OnMain() {
	fmt.Println("Starting up")
	for i, runner := range Servers {
		fmt.Printf("%v. Starting %s\n", i, runtime.FuncForPC(reflect.ValueOf(runner).Pointer()).Name())
		go runner()
	}

	loader.RunningForever()
}