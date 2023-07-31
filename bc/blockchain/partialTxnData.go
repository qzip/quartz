/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package blockchain

/*
Partial Transaction Data.
*/

// HashedTuple hashed Key / Value pairs
type HashedTuple struct {
	Key []byte `json:"key"`
	Val []byte `json:"value"`
}

//PartialTxnData partial data
type PartialTxnData struct {
	TxnBlock    []BlockTxn             `json:"txnBlock"`
	Height      int64                  `json:"height"`
	MetaData    map[string]interface{} `json:"metadata,omitempty"`
	PrivateData []HashedTuple          `json:"private,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}
