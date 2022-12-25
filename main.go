package main

import (
	"lets-go-framework/bootstraps"
	"lets-go-framework/controllers"
	"lets-go-framework/initiators"
)

// Initialize all required vars, consts
func init() {
	bootstraps.OnInit()
}

func main() {
	bootstraps.OnMain()

	controllers.CreateOrder()

	initiators.RunningForever()
}
