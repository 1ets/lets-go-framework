package lets

import (
	"fmt"
	"reflect"
	"runtime"
)

type Bootstrap struct {
	OnInits []func()
	OnMains []func()
}

func (b *Bootstrap) OnInit() {
	fmt.Println("Initialization")
	for i, initializer := range b.OnInits {
		fmt.Printf("%v. Initializing %s\n", i, runtime.FuncForPC(reflect.ValueOf(initializer).Pointer()).Name())
		initializer()
	}
}

func (b *Bootstrap) OnMain() {
	fmt.Println("Starting up")
	for i, runner := range b.OnMains {
		fmt.Printf("%v. Starting %s", i, runtime.FuncForPC(reflect.ValueOf(runner).Pointer()).Name())
		go runner()
	}

	runningForever()
}
