package types

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

// Default configuration
const (
	SERVE_HTTP_PORT = "5000"
)

// Interface for accessable method
type IHttpServer interface {
	GetPort() string
	GetMiddleware() func(*gin.Engine)
	GetRouter() func(*gin.Engine)
}

// Serve information
type HttpServer struct {
	Port       string
	Middleware func(*gin.Engine)
	Router     func(*gin.Engine)
}

// Get Port
func (hs *HttpServer) GetPort() string {
	if hs.Port == "" {
		golog.Warn("Configs Http: SERVE_HTTP_PORT is not set, using default configuration.")

		return SERVE_HTTP_PORT
	}

	return hs.Port
}

// Get Middleware
func (hs *HttpServer) GetMiddleware() func(*gin.Engine) {
	return hs.Middleware
}

// Get Router
func (hs *HttpServer) GetRouter() func(*gin.Engine) {
	return hs.Router
}
