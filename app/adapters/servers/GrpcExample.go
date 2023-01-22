package servers

import (
	"context"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/adapters/protobuf"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"
	"net/http"
)

type GrpcExample struct {
	protobuf.ExampleServiceServer
}

func (GrpcExample) Example(ctx context.Context, rx *protobuf.RequestExample) (tx *protobuf.ResponseExample, err error) {
	// Data binding
	var request data.RequestExample
	lets.Bind(rx, &request)

	lets.LogD("gRPC Server: receive \n%s", lets.ToJson(request))

	// Call controller
	response, err := controllers.GrpcServerExample(request)
	if err != nil {
		lets.LogE("gRPC Server: GrpcExample.Greeting:", err.Error())
		return
	}

	var greeting protobuf.Greeting
	lets.Bind(response, &greeting)

	// Convert response data type to protobuf type
	tx = &protobuf.ResponseExample{
		Code:   http.StatusOK,
		Status: "success",
		Data:   &greeting,
	}

	return
}

// func (GrpcUser) Find(ctx context.Context, request *protobuf.RequestUserFind) (response *protobuf.ResponseUserFind, err error) {
// 	fmt.Println("ApiUserServer.Find()")

// 	// Convert protobuf type to data type
// 	requestJson, _ := json.Marshal(request)
// 	fmt.Println(string(requestJson))

// 	// Data binding
// 	var cRequest data.RequestUserFind
// 	json.Unmarshal(requestJson, &cRequest)

// 	// Call controller
// 	// controllerResponse, err := controllers.UserFind(&cRequest)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())

// 	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 	// 		// Convert response data type to protobuf type
// 	// 		response = &protobuf.ResponseUserFind{
// 	// 			Code:   http.StatusNotFound,
// 	// 			Status: "success",
// 	// 		}
// 	// 		err = nil // reset not error
// 	// 	}

// 	// 	return
// 	// }

// 	// fmt.Printf("Found: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

// 	// // Convert response data type to protobuf type
// 	// response = &protobuf.ResponseUserFind{
// 	// 	Code:   http.StatusFound,
// 	// 	Status: "success",
// 	// 	Data: &protobuf.User{
// 	// 		Id:      uint64(controllerResponse.ID),
// 	// 		Name:    controllerResponse.Name,
// 	// 		Balance: float32(controllerResponse.Balance),
// 	// 	},
// 	// }

// 	return
// }

// func (GrpcUser) Get(ctx context.Context, request *protobuf.RequestUserGet) (response *protobuf.ResponseUserGet, err error) {
// 	fmt.Println("ServiceUser.Get()")
// 	// // Call controller
// 	// controllerResponse, err := controllers.UserGet(&data.RequestUserGet{})
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// response = &protobuf.ResponseUserGet{
// 	// 	Code:   200,
// 	// 	Status: "success",
// 	// }

// 	// for _, account := range controllerResponse {
// 	// 	response.Data = append(response.Data, &protobuf.User{
// 	// 		Id:      uint64(account.ID),
// 	// 		Name:    account.Name,
// 	// 		Balance: float32(account.Balance),
// 	// 	})
// 	// }
// 	return
// }

// func (GrpcUser) Update(ctx context.Context, request *protobuf.RequestUserUpdate) (response *protobuf.ResponseUserUpdate, err error) {
// 	fmt.Println("ServiceUser.Update()")

// 	// // Data binding
// 	// cRequest := data.RequestUserUpdate{
// 	// 	Id: uint(request.GetFind().GetId()),
// 	// 	Fields: struct {
// 	// 		Name    string  "json:\"name,omitempty\""
// 	// 		Balance float64 "json:\"balance,omitempty\""
// 	// 	}{
// 	// 		Name:    request.GetData().GetName(),
// 	// 		Balance: float64(request.GetData().GetBalance()),
// 	// 	},
// 	// }

// 	// // Call controller
// 	// controllerResponse, err := controllers.UserUpdate(&cRequest)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// fmt.Printf("Updated: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

// 	// // Convert response data type to protobuf type
// 	// response = &protobuf.ResponseUserUpdate{
// 	// 	Code:   http.StatusAccepted,
// 	// 	Status: "success",
// 	// }
// 	return
// }

// func (GrpcUser) Delete(ctx context.Context, request *protobuf.RequestUserDelete) (response *protobuf.ResponseUserDelete, err error) {
// 	fmt.Println("ServiceUser.Delete()")

// 	// // Data binding
// 	// cRequest := data.RequestUserDelete{
// 	// 	Id: uint(request.GetFind().GetId()),
// 	// }

// 	// // Call controller
// 	// controllerResponse, err := controllers.UserDelete(&cRequest)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	return
// 	// }

// 	// fmt.Printf("Updated: %s balance %2.f \n", controllerResponse.Name, controllerResponse.Balance)

// 	// Convert response data type to protobuf type
// 	response = &protobuf.ResponseUserDelete{
// 		Code:   http.StatusAccepted,
// 		Status: "success",
// 	}
// 	return
// }
