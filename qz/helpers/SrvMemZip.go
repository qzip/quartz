package helpers

import (
	"context"
	"qz/zipfs/srvzip"
	"reflect"
)

// SrvMemZip implements qz/commands.HelperFactory
type SrvMemZip struct {
}

// Name implements qz/commands.HelperFactory
func (fs *SrvMemZip) Name() string {
	return reflect.TypeOf(fs).Elem().String()
}

// Help implements qz/commands.HelperFactory
func (fs *SrvMemZip) Help() string {
	return `
	"component": "qz/helpers.SrvMemZip",
	"param": ""
	`
}

// ComponentType implements qz/commands.HelperFactory
func (fs *SrvMemZip) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

// B64MemBlock returns a B64 encoded string
type B64MemBlock interface {
	GetMemBlock() string
}

// CreateHelper returns http.Handler implements qz/commands.HelperFactory.
func (fs *SrvMemZip) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (hnd interface{}, err error) {
	b64 := srvzip.MemZippedB64
	if param != nil {
		p, ok := param.(string)
		if ok {
			if hnd := ctx.Value(p); hnd != nil {
				if memb, ok := hnd.(B64MemBlock); ok {
					b64 = memb.GetMemBlock()
				}
			}
		}
	}
	zs, err := srvzip.NewMemZipServer(b64)
	if err != nil {
		return
	}
	hnd = zs.MemZipServer()
	return
}
