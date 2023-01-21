package lets

import (
	"github.com/kataras/golog"
)

var (
	Log  = golog.New()
	LogD = Log.Debugf
	LogI = Log.Infof
	LogW = Log.Warnf
	LogE = Log.Errorf
	LogF = Log.Fatalf
	Logf = Log.Logf
)
