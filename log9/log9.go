package log9

import (
	"github.com/kataras/golog"
)

// Print debug message
func D(tag string, data interface{}) {
	var format string

	switch data.(type) {
	case int:
		format = "%s: (%T) %v"
	case float64:
		format = "%s: (%T) %f"
	case string:
		format = "%s: (%T) %s"
	default:
		golog.Errorf("%s: (%T) %s")
	}

	print("d", format, tag, data)
}

// Print to terminal
func print(outType string, format string, tag string, data interface{}) {
	switch outType {
	case "d":
		golog.Errorf(format, tag, data, data)
	case "i":
		golog.Errorf(format, tag, data, data)
	case "e":
		golog.Errorf(format, tag, data, data)
	case "f":
		golog.Errorf(format, tag, data, data)
	default:
		golog.Error("Invalid Format")
	}
}

// TODO: support using / not using data type
// TODO: config formatting
