package boot

import (
	"fmt"
	"lets-go-framework/lets/frameworks"
	"lets-go-framework/lets/loader"
	"reflect"
	"runtime"
)

// List of initializer
var Initializer = []func(){
	loader.Environment,
}

// List of framework that start on lets
var Servers = []func(){
	frameworks.Http,
	// frameworks.Grpc,
	// frameworks.RabbitMQ,
}

// Add initialization function and run before application starting
func AddInitializer(init func()) {
	Initializer = append(Initializer, init)
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
