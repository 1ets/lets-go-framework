package controllers

import (
	"fmt"
	"lets-go-framework/lets"
)

func RabbitTransferResult(r lets.MessageContext) {
	fmt.Println(string(r.GetData()))
}
