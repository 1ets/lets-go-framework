package lets

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type httpResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func HttpResponseJson(g *gin.Context, i interface{}, err error) {
	if err != nil {
		g.JSON(500, httpResponse{
			Status:  "error",
			Message: status.Convert(err).Message(),
		})
		return
	}

	statusCode := http.StatusOK
	status := "success"

	var data map[string]interface{}
	dataJson, _ := json.Marshal(i)
	json.Unmarshal(dataJson, &data)

	if data["code"] != nil && data["code"] != 0 {
		statusCode = int(data["code"].(float64))
	}

	if data["status"] != nil && data["status"] != "" {
		status = data["status"].(string)
	}

	delete(data, "code")
	delete(data, "status")

	g.JSON(statusCode, httpResponse{
		Status: status,
		Data:   data,
	})
}
