package services

import (
	"fmt"
	"lets-go-framework/app/adapters/servers"
	"time"

	"github.com/gin-gonic/gin"
)

// Global Middleware setup
func HttpMiddleware(middleware *gin.Engine) {
	var defaultLogFormatter = func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		return fmt.Sprintf("%s[HTTP]%s %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			"\x1b[32m", resetColor,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}

	middleware.Use(gin.LoggerWithFormatter(defaultLogFormatter))
	middleware.Use(gin.Recovery())
}

func HttpRouter(route *gin.Engine) {
	route.POST("example", servers.HttpPostExample)
}
