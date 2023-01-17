package loader

import (
	"math"
	"time"
)

// Hold the thread for exitting
func RunningForever() {
	for {
		time.Sleep(time.Duration(math.MaxInt64))
	}
}

// TODO: Create stopper
// TODO: Fatal handling
