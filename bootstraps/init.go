package bootstraps

import (
	"lets-go-framework/configs"
)

// Run bootstrap when on init function running
func OnInit() {
	LoadEnv()
}

func OnMain() {

	configs.Initialize()
}
