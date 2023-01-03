package lets

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
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
			Message: status.Convert(err).Message(),
		})
		return
	}

	statusCode := http.StatusOK
	status := "success"
	if i != nil {
		var data map[string]interface{}
		dataJson, _ := json.Marshal(i)
		json.Unmarshal(dataJson, &data)

		if data["code"] != nil {
			statusCode = int(data["code"].(float64))
			i = nil
		}

		if data["status"] != nil {
			status = data["status"].(string)
			i = nil
		}
	}

	g.JSON(statusCode, HttpResponse{
		Status: status,
		Data:   i,
	})
}
