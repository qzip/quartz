package imle

import "time"

//https://blog.ethereum.org/2015/11/15/merkling-in-ethereum/

//Block Bitcoin like Block
type Block struct {
	PreviousHash []byte         `json:"previousHash"`
	Nonce        []byte         `json:"nonce"`
	Timestamp    time.Time      `json:"tmstamp"`
	MerkleRoot   []byte         `json:"merkleRoot"`
	TxnLog       *MerkleDataLog `json:"txnLog"`
}
