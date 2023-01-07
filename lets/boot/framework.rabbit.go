package boot

import (
	"fmt"
	"lets-go-framework/adapters"
	"lets-go-framework/configs"
	"lets-go-framework/lets/drivers"
	"lets-go-framework/services"
)

// Define rabbit service host and port
func LoadRabbitFramework() {
	fmt.Println("LoadRabbitFramework()")

	rabbit := drivers.ServiceRabbit{
		Dsn:      configs.RabbitDsn,
		Consumer: configs.RabbitConsumer,
		Engine:   drivers.MessageEngine{},
	}

	rabbit.Init()

	services.RabbitEventHandler(&rabbit.Engine)
	adapters.RabbitRegister(&rabbit)

	var err error
	err = rabbit.Connect()
	if err != nil {
		fmt.Printf("ERROR rabbit.Serve(): %s\n", err.Error())
		return
	}

	err = rabbit.Register()
	if err != nil {
		fmt.Printf("ERROR rabbit.Register(): %s\n", err.Error())
		return
	}
}
