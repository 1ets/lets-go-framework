package lets

import (
	"lets-go-framework/configs"
)

type Bootstrap struct{}

func (b *Bootstrap) OnInit() {
	loadEnv()
}

func (b *Bootstrap) OnMain() {
	configs.Initialize()

	go loadHttpFramework()
	go loadGrpcFramework()

	runningForever()
}
