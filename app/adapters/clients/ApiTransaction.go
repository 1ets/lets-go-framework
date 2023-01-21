package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/adapters/protobuf"
)

var ApiTransaction = apiTransaction{}

type apiTransaction struct {
	Client protobuf.TransactionServiceClient
}

func (ac *apiTransaction) Get(c context.Context, data data.RequestGetTransaction) (*protobuf.ResponseGetTransaction, error) {
	fmt.Println("GetTransaction")

	adapt, _ := json.Marshal(data)

	var message = protobuf.RequestGetTransaction{}
	json.Unmarshal(adapt, &message)

	return ac.Client.Get(context.Background(), &message)
}
