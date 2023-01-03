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
	example.GET("account", controllers.HttpAccountGet)
	example.GET("account/:id", controllers.HttpAccountFind)
	example.POST("account", controllers.HttpAccountRegister)

	// Transfers
	example.GET("transfer", controllers.HttpAccountGet)
	example.POST("transfer/success", controllers.HttpAccountGet)
}
