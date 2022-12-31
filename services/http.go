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
	example.GET("account/:id", controllers.PostTransferMoney)
	example.POST("account", controllers.PostTransferMoney)

	// Transfers
	example.GET("transfer", controllers.PostTransferMoney)
	example.POST("transfer/success", controllers.PostTransferMoney)
}
