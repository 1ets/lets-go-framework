package controllers

import (
	"fmt"
	"lets-go-framework/lets/drivers"
)

func RabbitTransferResult(r drivers.MessageContext) {
	fmt.Println(string(r.GetData()))
}
