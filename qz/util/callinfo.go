package util

import (
	"context"
	"fmt"
	"time"
)

// DebugInfoCtxKey key to context
var DebugInfoCtxKey = "debug"

const DebugHandlerKey = "debugHandler"

// Debug Global setting
var Debug = true

// DebugInfoHandler component must implement this interface
type DebugInfoHandler interface {
	Print(ctx context.Context, msg string)
}

// DebugInfo : Call Debugger component if present
func DebugInfo(ctx context.Context, msg string) {
	if Debug {
		if handler, ok := ctx.Value(DebugHandlerKey).(DebugInfoHandler); ok {
			handler.Print(ctx, msg)
		} else {
			fmt.Println(time.Now(), msg)
		}
	}
}

// SetDebugFlag set debug if "debug" is true in config
func SetDebugFlag(cfg map[string]interface{}) bool {
	v, ok := cfg[DebugInfoCtxKey]
	if v != nil && ok {
		b, ok := v.(bool)
		if ok {
			Debug = b
		}
	}
	return Debug
}
