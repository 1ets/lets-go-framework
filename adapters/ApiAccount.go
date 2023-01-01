package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/adapters/protobuf"
)

// GRPC API collections for Account Service
var ApiAccount = apiAccountAdapter{}

// Adapter for GRPC API Account Service
type apiAccountAdapter struct {
	Client protobuf.AccountServiceClient
}

// API Collection: Get accounts
func (ac *apiAccountAdapter) Get(ctx context.Context, request *data.RequestGetAccount) (response *data.ResponseGetAccount, err error) {
	fmt.Println("GetAccount")

	// Convert data type to protobuf type
	reqGetAccount, _ := json.Marshal(request)

	var getAccount = protobuf.RequestGetAccount{}
	json.Unmarshal(reqGetAccount, &getAccount)

	// Start call endpoint
	apiResponse, err := ac.Client.Get(context.Background(), &getAccount)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Convert response protobuf type to data type
	response = &data.ResponseGetAccount{
		Id:      uint(apiResponse.Id),
		Name:    apiResponse.Name,
		Balance: float64(apiResponse.Balance),
	}

	return
}
