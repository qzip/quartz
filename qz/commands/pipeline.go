/*
Copyright (c) 2024, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package commands

import "context"

type Pipeline interface {
	Process(ctx context.Context)
}

type Source interface {
	Pipeline
	Chan(ctx context.Context, errCh chan error, cancel context.CancelFunc) chan interface{}
}

type Sink interface {
	Pipeline
	Chan(ctx context.Context, source chan interface{}, errCh chan error, cancel context.CancelFunc)
}

type Transformer interface {
	Pipeline
	Chan(ctx context.Context, source chan interface{}, errCh chan error, cancel context.CancelFunc) (sink chan interface{})
}
