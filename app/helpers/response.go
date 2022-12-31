package helpers

import (
	"lets-go-framework/app/structs"

	"github.com/gin-gonic/gin"
)

func GinResponse(g *gin.Context, i interface{}, err error) {
	if err != nil {
		g.JSON(500, structs.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	g.JSON(200, structs.Response{
		Status: "success",
		Data:   i,
	})
}
