package services

import (
	"lets-go-framework/app/controllers"

	"github.com/gin-gonic/gin"
)

// Global Middleware setup
func MiddlewareHttpService(middleware *gin.Engine) {
	middleware.Use(gin.Logger(), gin.Recovery())
}

func RouteHttpServicse(route *gin.Engine) {
	example := route.Group("example")

	// Accounts
	example.GET("account", controllers.HttpGetAccount)
	example.GET("account/:id", controllers.HttpGetAccount)
	example.POST("account", controllers.HttpGetAccount)

	// Transfers
	example.GET("transfer", controllers.HttpGetAccount)
	example.POST("transfer/success", controllers.HttpGetAccount)
}
