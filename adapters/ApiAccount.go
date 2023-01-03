package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/adapters/protobuf"
)

// GRPC API collections for Account Service
var ApiAccount = adapterGrpcAccount{}

// Adapter for GRPC API Account Service
type adapterGrpcAccount struct {
	Client protobuf.ApiAccountClient
}

// API Collection: Get accounts
func (ac *adapterGrpcAccount) Get(ctx context.Context, request *data.RequestGetAccount) (response *data.ResponseAccountGet, err error) {
	fmt.Println("GetAccount")

	// Convert data type to protobuf type
	requestJson, _ := json.Marshal(request)

	var requestGrpc = protobuf.RequestAccountGet{}
	json.Unmarshal(requestJson, &requestGrpc)

	// Call endpoint
	grpcResponse, err := ac.Client.Get(context.Background(), &requestGrpc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Convert response protobuf type to data type
	var responseBucket data.ResponseAccountGet
	for _, account := range grpcResponse.GetData() {
		responseBucket = append(responseBucket, data.Account{
			ID:      uint(account.Id),
			Name:    account.Name,
			Balance: float64(account.Balance),
		})
	}

	response = &responseBucket

	return
}
