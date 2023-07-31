package defs

import (
	"context"
	bc "qz/blockchain"
	"qz/cas"
)

// Exporter helper to export the blockchain
type Exporter interface {
	ChainHead(ctx context.Context) (bc *bc.BlockChain, err error)
	BlockChain(ctx context.Context, blk int64) (*bc.BlockChain, error)
	GetCAS(ctx context.Context, casHash []byte) (*cas.CasJSON, error)
	Signatures(ctx context.Context, txnHash []byte) ([]bc.Signature, error)
}
