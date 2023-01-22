package app

import (
	"lets-go-framework/config"
	"lets-go-framework/lets/boot"
)

// Intercept lets initialization
func OnInit() {
	boot.AddInitializer(config.App)
	boot.AddInitializer(config.Database)
}
