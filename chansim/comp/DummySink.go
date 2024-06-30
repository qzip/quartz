package comp

import (
	"context"
	"fmt"
	"qz/util"
	"reflect"
)

type DummySink struct {
	source     chan interface{}
	errCh      chan error
	cancel     context.CancelFunc
	count      int
	chanInType string
}

// Process implements commands.Pipeline
func (dbf *DummySink) Process(ctx context.Context) {
	for in := range dbf.source {
		select {
		case <-ctx.Done():
			dbf.errCh <- fmt.Errorf("%v: context Done", reflect.TypeOf(*dbf).String())
			return
		default:
			dbf.chanInType = reflect.TypeOf(in).String()
			dbf.count++
		}
	}
	util.DebugInfo(ctx, fmt.Sprintf("DummySource.Process done count %v, chan input type [%v]", dbf.count, dbf.chanInType))

}

// Chan implements commands.Sink for qz/cmd.ChanExec
func (dbf *DummySink) Chan(ctx context.Context, source chan interface{}, errCh chan error, cancel context.CancelFunc) {
	dbf.source = source
	dbf.errCh = errCh
	dbf.cancel = cancel

}
