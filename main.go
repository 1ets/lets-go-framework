package main

import "lets-go-framework/lets"

// Initialize all required vars, consts

var boot = lets.Bootstrap{}

func init() {
	boot.OnInit()
}

func main() {
	boot.OnMain()
}
