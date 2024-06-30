package comp

import (
	"context"
	"fmt"
	"qz/util"
	"reflect"
	"time"
)

type DummyTransformer struct {
	count      int
	lastMsg    time.Time
	errChan    chan error
	cancel     context.CancelFunc
	downstream chan interface{}
	upstream   chan interface{}
}

func (dcs *DummyTransformer) Chan(ctx context.Context, source chan interface{}, errCh chan error, cancel context.CancelFunc) (sink chan interface{}) {
	dcs.cancel = cancel
	dcs.errChan = errCh
	dcs.upstream = source
	dcs.downstream = make(chan interface{})

	return dcs.downstream
}
func (dcs *DummyTransformer) Process(ctx context.Context) {
	for in := range dcs.upstream {
		select {
		case <-ctx.Done():
			dcs.error(fmt.Errorf("%v: context Done", reflect.TypeOf(dcs).String()))
			return
		default:
			dcs.downstream <- in
			dcs.count++
			dcs.lastMsg = time.Now()
		}

	}
	close(dcs.downstream)
	util.DebugInfo(ctx, fmt.Sprintf("DummyTransformer.Process done count %v, last %v", dcs.count, dcs.lastMsg))
}

func (dcs *DummyTransformer) error(err error) {
	dcs.errChan <- err
	close(dcs.downstream)
	dcs.cancel()
}
