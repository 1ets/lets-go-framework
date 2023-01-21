package frameworks

import (
	"fmt"
	"lets-go-framework/lets/types"
	"time"

	"github.com/gin-gonic/gin"
)

var HttpConfig types.IHttpServer

// HTTP service struct
type httpService struct {
	Server     string
	Engine     *gin.Engine
	Middleware func(*gin.Engine)
	Router     func(*gin.Engine)
}

// Initialize service
func (http *httpService) Init() {
	fmt.Println("httpService.Init()")

	http.Server = fmt.Sprintf(":%s", HttpConfig.GetPort())
	http.Engine = gin.New()
	http.Middleware = HttpConfig.GetMiddleware()
	http.Router = HttpConfig.GetRouter()

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

	http.Engine.Use(gin.LoggerWithFormatter(defaultLogFormatter))
}

// Run service
func (http *httpService) Serve() {
	fmt.Println("httpService.Serve()")

	http.Engine.Run(http.Server)
}

// Start http service
func Http() {
	if HttpConfig == nil {
		return
	}

	fmt.Println("httpService.LoadHttpFramework()")

	var http httpService

	http.Init()
	http.Middleware(http.Engine)
	http.Router(http.Engine)
	http.Serve()
}
