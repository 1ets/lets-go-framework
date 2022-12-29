package services

import (
	"lets-go-framework/controllers"

	"github.com/gin-gonic/gin"
)

// Global Middleware setup
func MiddlewareHttpService(middleware *gin.Engine) {
	middleware.Use(gin.Logger(), gin.Recovery())
}

func RouteHttpServicse(route *gin.Engine) {
	example := route.Group("example")
	example.GET("transfer/success", controllers.PostTransferMoney)
	example.GET("transfer/failed", controllers.PostTransferMoneyFailed)
}
