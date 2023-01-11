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
	example.PATCH("account/:id", controllers.HttpAccountUpdate)
	example.DELETE("account/:id", controllers.HttpAccountRemove)

	// Transfers
	example.GET("transaction", controllers.HttpTransactionGet)
	example.POST("transfer-stateful", controllers.HttpTransferStatefull)
	example.POST("transfer-stateless", controllers.HttpTransferStateless)
}
