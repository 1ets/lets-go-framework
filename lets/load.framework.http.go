package lets

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
	http.Server = gin.New()
}

func (http *httpService) Serve() {
	ServePort := fmt.Sprintf(":%s", http.Port)
	http.Server.Run(ServePort)
}

// Define http service host and port
func loadHttpFramework() {
	http := httpService{
		Port: os.Getenv("SERVE_HTTP_PORT"),
	}

	http.Init()

	services.MiddlewareHttpService(http.Server)
	services.RouteHttpServicse(http.Server)

	http.Serve()
}
