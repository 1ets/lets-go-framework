package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/adapters/protobuf"
)

var ApiAccount = apiAccount{}

type apiAccount struct {
	Client protobuf.AccountServiceClient
}

func (ac *apiAccount) Get(c context.Context, data data.RequestGetAccount) (*protobuf.ResponseGetAccount, error) {
	fmt.Println("GetAccount")

	reqGetAccount, _ := json.Marshal(data)

	var getAccount = protobuf.RequestGetAccount{}
	json.Unmarshal(reqGetAccount, &getAccount)

	return ac.Client.Get(context.Background(), &getAccount)
}
