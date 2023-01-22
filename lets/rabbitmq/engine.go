package rabbitmq

import (
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"reflect"
	"runtime"
)

// Engine for controller
type Engine struct {
	handlers map[string]func(*types.Event)
	Debug    bool
}

func (me *Engine) Event(name string, handler func(*types.Event)) {
	if me.Debug {
		lets.LogD("Rabbit Event: %-20s --> %v", name, runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
	}

	if me.handlers == nil {
		me.handlers = map[string]func(*types.Event){}
	}

	me.handlers[name] = handler
}

func (me *Engine) Call(name string, event *types.Event) {
	if adapter := me.handlers[name]; adapter != nil {
		if me.Debug {
			lets.LogD("Rabbit Event: %-20s --> %v", name, runtime.FuncForPC(reflect.ValueOf(adapter).Pointer()).Name())
		}
		adapter(event)
		return
	}

	lets.LogE("Rabbit Event: not found: %s", name)
}
