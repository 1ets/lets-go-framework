package data

import "lets-go-framework/lets/types"

type RequestExample struct {
	Name string `json:"name"`
}

type ResponseExample struct {
	types.Response

	Greeting string `json:"greeting"`
}

type RequestDatabaseExample struct{}

type ResponsDatabaseeExample struct {
	types.Response

	Note string `json:"note"`
}

// type RequestGrpcExample struct {
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }

// type ResponseGrpcExample struct {
// 	types.Response

// 	Data User `json:"data"`
// }

// type RequestRabbitExample struct {
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }

// type ResponsRabbitcExample struct {
// 	types.Response

// 	Data User `json:"data"`
// }
