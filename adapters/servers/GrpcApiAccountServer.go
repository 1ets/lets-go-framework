package servers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/adapters/protobuf"
	"lets-go-framework/app/controllers"
	"net/http"

	"gorm.io/gorm"
)

type ApiAccountServer struct {
	protobuf.ApiAccountServer
}

func (ApiAccountServer) Insert(ctx context.Context, request *protobuf.RequestAccountInsert) (response *protobuf.ResponseAccountInsert, err error) {
	fmt.Println("ApiAccountServer.Insert()")

	// Convert protobuf type to data type
	requestJson, _ := json.Marshal(request)

	// Data binding
	var cRequest data.RequestAccountInsert
	json.Unmarshal(requestJson, &cRequest)

	// Call controller
	controllerResponse, err := controllers.AccountInsert(&cRequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Created: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

	// Convert response data type to protobuf type
	response = &protobuf.ResponseAccountInsert{
		Code:   http.StatusCreated,
		Status: "success",
	}

	return
}

func (ApiAccountServer) Find(ctx context.Context, request *protobuf.RequestAccountFind) (response *protobuf.ResponseAccountFind, err error) {
	fmt.Println("ApiAccountServer.Find()")

	// Convert protobuf type to data type
	requestJson, _ := json.Marshal(request)
	fmt.Println(string(requestJson))

	// Data binding
	var cRequest data.RequestAccountFind
	json.Unmarshal(requestJson, &cRequest)

	// Call controller
	controllerResponse, err := controllers.AccountFind(&cRequest)
	if err != nil {
		fmt.Println(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Convert response data type to protobuf type
			response = &protobuf.ResponseAccountFind{
				Code:   http.StatusNotFound,
				Status: "success",
			}
			err = nil // reset not error
		}

		return
	}

	fmt.Printf("Found: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

	// Convert response data type to protobuf type
	response = &protobuf.ResponseAccountFind{
		Code:   http.StatusFound,
		Status: "success",
		Data: &protobuf.Account{
			Id:      uint64(controllerResponse.ID),
			Name:    controllerResponse.Name,
			Balance: float32(controllerResponse.Balance),
		},
	}

	return
}

func (ApiAccountServer) Get(ctx context.Context, request *protobuf.RequestAccountGet) (response *protobuf.ResponseAccountGet, err error) {
	fmt.Println("ServiceAccount.Get()")
	// Call controller
	controllerResponse, err := controllers.AccountGet(&data.RequestAccountGet{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response = &protobuf.ResponseAccountGet{
		Code:   200,
		Status: "success",
	}

	for _, account := range controllerResponse {
		response.Data = append(response.Data, &protobuf.Account{
			Id:      uint64(account.ID),
			Name:    account.Name,
			Balance: float32(account.Balance),
		})
	}
	return
}

func (ApiAccountServer) Update(ctx context.Context, request *protobuf.RequestAccountUpdate) (response *protobuf.ResponseAccountUpdate, err error) {
	fmt.Println("ServiceAccount.Update()")

	// Data binding
	cRequest := data.RequestAccountUpdate{
		Id: uint(request.GetFind().GetId()),
		Fields: struct {
			Name    string  "json:\"name,omitempty\""
			Balance float64 "json:\"balance,omitempty\""
		}{
			Name:    request.GetData().GetName(),
			Balance: float64(request.GetData().GetBalance()),
		},
	}

	// Call controller
	controllerResponse, err := controllers.AccountUpdate(&cRequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Updated: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

	// Convert response data type to protobuf type
	response = &protobuf.ResponseAccountUpdate{
		Code:   http.StatusAccepted,
		Status: "success",
	}
	return
}
func (ApiAccountServer) Delete(ctx context.Context, request *protobuf.RequestAccountDelete) (response *protobuf.ResponseAccountDelete, err error) {
	fmt.Println("ServiceAccount.Delete()")
	fmt.Println("ServiceAccount.Update()")

	// Data binding
	cRequest := data.RequestAccountDelete{
		Id: uint(request.GetFind().GetId()),
	}

	// Call controller
	controllerResponse, err := controllers.AccountDelete(&cRequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Updated: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

	// Convert response data type to protobuf type
	response = &protobuf.ResponseAccountDelete{
		Code:   http.StatusAccepted,
		Status: "success",
	}
	return
}
