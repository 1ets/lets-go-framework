package lets

import (
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(g *gin.Context, i interface{}, err error) {
	if err != nil {
		g.JSON(500, HttpResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	g.JSON(200, HttpResponse{
		Status: "success",
		Data:   i,
	})
}
