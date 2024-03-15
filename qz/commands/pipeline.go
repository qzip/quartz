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
	Chan(ctx context.Context) (chan interface{}, error)
}

type Sink interface {
	Pipeline
	Chan(ctx context.Context, source chan interface{}) error
}

type Transformer interface {
	Pipeline
	Chan(ctx context.Context, source chan interface{}) (sink chan interface{}, err error)
}
