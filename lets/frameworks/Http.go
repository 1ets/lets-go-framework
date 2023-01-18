package frameworks

import (
	"fmt"
	"lets-go-framework/lets/types"

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
}

// Run service
func (http *httpService) Serve() {
	fmt.Println("httpService.Serve()")

	http.Engine.Run(http.Server)
}

// Start http service
func Http() {
	fmt.Println("httpService.LoadHttpFramework()")

	var http httpService

	http.Init()
	http.Middleware(http.Engine)
	http.Router(http.Engine)
	http.Serve()
}
