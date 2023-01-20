package main

import (
	"lets-go-framework/app"
	"lets-go-framework/lets/boot"
)

// Initiate lets engine
func init() {
	app.OnInit()
	boot.OnInit()
}

// Bootstrap applications
func main() {
	boot.OnMain()
}
