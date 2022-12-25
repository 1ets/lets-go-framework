package initiators

import (
	"lets-go-framework/log9"
	"math"
	"time"
)

// Hold the thread for exitting
func RunningForever() {
	for {
		log9.D("RunningForever", "Service never stopped")
		time.Sleep(time.Duration(math.MaxInt64))
	}
}

// TODO: Create stopper
// TODO: Fatal handling
