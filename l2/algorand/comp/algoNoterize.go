package comp

import (
	"context"
	"qz/seq"
)

const (
	AlgodAddress = "http://localhost"
	AlgodPort    = "4001"
)

type AlgoNotarize struct {
	Base64PrivateKey string `json:"base64PrivateKey"`
	AlgodAddress     string `json:"AlgodAddress,omitempty"`
	AlgodPort        string `json:"AlgodPort,omitempty"`

	//private fields
	helper  seq.CtxHelper
	errChan chan error
}

func (an *AlgoNotarize) SetCtxHelper(hlp seq.CtxHelper) {
	an.helper = hlp
}

func (an *AlgoNotarize) SetChanErr(ce chan error) {
	an.errChan = ce
}

//https://github.com/algorand/go-algorand-sdk/blob/v2.4.0/examples/utils.go

func (an *AlgoNotarize) Process(ctx context.Context) {
}
