package bc

import "time"

type ChainInfo struct {
	Flavour string `json:"flavour"` // eg. algorand,burrow
	ChainID string `json:"chainID"`
	Height  int64  `json:"height"`
	Hash    string `json:"hash"`
	TxnID   string `json:"txnID"`
}
type Peg struct {
	MainChain ChainInfo `json:"mainChain"`
	SideChain ChainInfo `json:"sideChain"`
	Tmstamp   time.Time `json:"tmstamp"`
}
