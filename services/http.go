package services

import (
	"fmt"
	"lets-go-framework/app/controllers"

	"github.com/gin-gonic/gin"
)

// Global Middleware setup
func MiddlewareHttpService(middleware *gin.Engine) {
	fmt.Println("MiddlewareHttpService()")
	middleware.Use(gin.Logger(), gin.Recovery())
}

func RouteHttpService(route *gin.Engine) {
	fmt.Println("RouteHttpService()")

	example := route.Group("example")

	// Accounts
	example.GET("account", controllers.HttpGetAccount)
	example.GET("account/:id", controllers.HttpGetAccount)
	example.POST("account", controllers.HttpGetAccount)

	// Transfers
	example.GET("transfer", controllers.HttpGetAccount)
	example.POST("transfer/success", controllers.HttpGetAccount)
}
