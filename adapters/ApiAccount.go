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
func (ac *adapterGrpcAccount) Get(ctx context.Context, request *data.RequestAccountGet) (response *data.ResponseAccountGet, err error) {
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
			Id:      uint(account.Id),
			Name:    account.Name,
			Balance: float64(account.Balance),
		})
	}

	response = &responseBucket

	return
}

// API Collection: Find account
func (ac *adapterGrpcAccount) Find(ctx context.Context, request *data.RequestAccountFind) (response *data.ResponseAccountFind, err error) {
	fmt.Println("Find Account")

	// Convert data type to protobuf type
	requestJson, _ := json.Marshal(request)

	var requestGrpc = protobuf.RequestAccountFind{}
	json.Unmarshal(requestJson, &requestGrpc)

	// Call endpoint
	grpcResponse, err := ac.Client.Find(context.Background(), &requestGrpc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Convert response protobuf type to data type
	payload := grpcResponse.GetData()
	response = &data.ResponseAccountFind{
		Id:      uint(payload.GetId()),
		Name:    payload.GetName(),
		Balance: float64(payload.GetBalance()),
	}

	return
}

// API Collection: Find account
func (ac *adapterGrpcAccount) Insert(ctx context.Context, request *data.RequestAccountInsert) (response *data.ResponseAccountInsert, err error) {
	fmt.Println("Insert Account")

	// Convert data type to protobuf type
	requestJson, _ := json.Marshal(request)

	var requestGrpc = protobuf.RequestAccountInsert{}
	json.Unmarshal(requestJson, &requestGrpc)

	// Call endpoint
	grpcResponse, err := ac.Client.Insert(context.Background(), &requestGrpc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Convert response protobuf type to data type
	response = &data.ResponseAccountInsert{
		Code:   uint16(grpcResponse.GetCode()),
		Status: grpcResponse.GetStatus(),
	}

	return
}
