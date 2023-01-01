package boot

import (
	"fmt"
	"lets-go-framework/services"
	"os"

	"github.com/gin-gonic/gin"
)

type httpService struct {
	Host   string
	Port   string
	Server *gin.Engine
}

func (http *httpService) Init() {
	fmt.Println("httpService.Init()")

	http.Server = gin.New()
}

func (http *httpService) Serve() {
	fmt.Println("httpService.Serve()")

	ServePort := fmt.Sprintf(":%s", http.Port)
	http.Server.Run(ServePort)
}

// Define http service host and port
func LoadHttpFramework() {
	fmt.Println("httpService.LoadHttpFramework()")

	http := httpService{
		Port: os.Getenv("SERVE_HTTP_PORT"),
	}

	http.Init()

	services.MiddlewareHttpService(http.Server)
	services.RouteHttpService(http.Server)

	http.Serve()
}
