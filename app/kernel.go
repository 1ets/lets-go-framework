package app

import (
	"lets-go-framework/config"
	"lets-go-framework/lets/boot"

	"github.com/gin-gonic/gin"
)

// Intercept lets initialization
func OnInit() {
	boot.AddInitializer(config.App)
	boot.AddInitializer(config.Database)

	// Set Gin
	gin.SetMode(gin.ReleaseMode)
}
