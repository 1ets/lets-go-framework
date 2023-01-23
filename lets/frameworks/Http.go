package frameworks

import (
	"fmt"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTP Configurations
var HttpConfig types.IHttpServer

// HTTP service struct
type httpServer struct {
	server     string
	engine     *gin.Engine
	middleware func(*gin.Engine)
	router     func(*gin.Engine)
}

// Initialize service
func (http *httpServer) init() {
	gin.SetMode(HttpConfig.GetMode())

	http.server = fmt.Sprintf(":%s", HttpConfig.GetPort())
	http.engine = gin.New()
	http.middleware = HttpConfig.GetMiddleware()
	http.router = HttpConfig.GetRouter()

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

	http.engine.Use(gin.LoggerWithFormatter(defaultLogFormatter))
}

// Run service
func (http *httpServer) serve() {
	http.engine.Run(http.server)
}

// Start http service
func Http() {
	if HttpConfig == nil {
		return
	}

	lets.LogI("HTTP Server Starting ...")

	var http httpServer

	http.init()
	http.middleware(http.engine)
	http.router(http.engine)
	http.serve()
}
