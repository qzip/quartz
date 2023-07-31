package defs

/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

import (
	"context"
	bc "qz/blockchain"
)

// Importer helper to save the records
// Go stacks do not limit the size, so passing by value in go routines should not be an issue.
// see: https://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite
type Importer interface {
	GetSavedBlockHeight(ctx context.Context) (int64, error)
	SaveBlockchain(ctx context.Context, bc bc.BlockChain) error
	SaveSignedData(ctx context.Context, signed bc.SignedData) error
}
