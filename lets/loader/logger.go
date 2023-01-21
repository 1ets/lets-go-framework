package loader

import (
	"encoding/json"
	"io"
	"lets-go-framework/lets"
	"sync"

	"github.com/kataras/golog"
	"github.com/kataras/pio"
)

var GinLevel golog.Level = 6

func Logger() {
	lets.Log.SetTimeFormat("2006-01-02 15:04:05")
	lets.Log.SetLevel("debug")
	// lets.Log.RegisterFormatter(&LetsFormatter{})
	// lets.Log.SetFormat("lets")

	golog.Levels[GinLevel] = &golog.LevelMetadata{
		Name:             "gin",
		AlternativeNames: []string{"http-server"},
		Title:            "[GIN]",
		ColorCode:        pio.Green,
	}
}

// LetsFormatter is a Formatter type for JSON logs.
type LetsFormatter struct {
	Indent string

	// Use one encoder per level, do not create new each time.
	// Even if the jser can set a different formatter for each level
	// on SetLevelFormat, the encoding's writers may be different
	// if that ^ is not called but SetLevelOutput did provide a different writer.
	encoders map[golog.Level]*json.Encoder
	mu       sync.RWMutex // encoders locker.
	encMu    sync.Mutex   // encode action locker.
}

// String returns the name of the Formatter.
// In this case it returns "json".
// It's used to map the formatter names with their implementations.
func (f *LetsFormatter) String() string {
	return "lets"
}

// Options sets the options for the JSON Formatter (currently only indent).
func (f *LetsFormatter) Options(opts ...interface{}) golog.Formatter {
	formatter := &LetsFormatter{
		Indent:   "  ",
		encoders: make(map[golog.Level]*json.Encoder, len(golog.Levels)),
	}

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if indent, ok := opt.(string); ok {
			formatter.Indent = indent
			break
		}
	}

	return formatter
}

// Format prints the logs in JSON format.
//
// Usage:
// logger.SetFormat("json") or
// logger.SetLevelFormat("info", "json")
func (f *LetsFormatter) Format(dest io.Writer, log *golog.Log) bool {
	f.mu.RLock()
	enc, ok := f.encoders[log.Level]
	f.mu.RUnlock()

	if !ok {
		enc = json.NewEncoder(dest)
		enc.SetIndent("", f.Indent)
		f.mu.Lock()
		f.encoders[log.Level] = enc
		f.mu.Unlock()
	}

	f.encMu.Lock()
	err := enc.Encode(log)
	f.encMu.Unlock()
	return err == nil
}
