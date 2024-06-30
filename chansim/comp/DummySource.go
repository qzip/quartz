package comp

import (
	"chansim"
	"context"
	"fmt"
	"qz/util"
	"time"
)

var MaxCount = 100

type DummySource struct {
	MaxCount   int
	count      int
	errChan    chan error
	cancel     context.CancelFunc
	downstream chan interface{}
}

// Process implements commands.Pipeline
func (dcs *DummySource) Process(ctx context.Context) {
	if dcs.MaxCount <= 0 {
		dcs.MaxCount = MaxCount
	}
	for dcs.count = 0; dcs.count < dcs.MaxCount; dcs.count++ {
		rec := &chansim.ChanSimPayload{Timestamp: time.Now(), Count: dcs.count}
		dcs.downstream <- rec
	}
	close(dcs.downstream)
	util.DebugInfo(ctx, fmt.Sprintf("DummySource.Process done count %v", dcs.count))

}

// Chan: implements commands.Source
func (dcs *DummySource) Chan(ctx context.Context, errCh chan error, cancel context.CancelFunc) chan interface{} {
	dcs.errChan = errCh
	dcs.cancel = cancel

	dcs.downstream = make(chan interface{})
	return dcs.downstream
}
