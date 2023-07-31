/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package blockchain

import (
	"bc/cas"
	"context"
	"crypto/rand"
	"encoding/hex"
	"merkle"
	"time"
)

/*
TODO: align with RFC 6962
https://tools.ietf.org/html/rfc6962

https://www.certificate-transparency.org/

https://github.com/google/certificate-transparency-go

https://gitlab.com/NebulousLabs/merkletree

https://en.wikipedia.org/wiki/Merkle_tree

https://github.com/google/trillian/

https://github.com/google/trillian/blob/master/docs/papers/VerifiableDataStructures.pdf


*/

// MerkleNode Block memory representation of block that is signed and placed in Blockchain record
type MerkleNode struct {
	TreeID   string       `json:"treeID"`
	Parent   cas.Hashable `json:"parent"`
	Nonce    string       `json:"nonce"`
	Tmstamp  time.Time    `json:"tmstamp"`
	TxnBlock []BlockTxn   `json:"txnBlock"`
}

//Hash implements merkle.Hasher interface
// computes merkle root of the BxnBlock param
func (mn *MerkleNode) Hash() []byte {
	return merkle.SimpleHashFromMap(mn.MerkleHasher())

}

//MerkleHasher transforms BlockTxn into merkle hasher map
func (mn *MerkleNode) MerkleHasher() map[string]merkle.Hasher {
	m := make(map[string]merkle.Hasher)
	m["treeID"] = cas.NewHashData([]byte(mn.TreeID))
	m["parent"] = mn.Parent
	m["nonce"] = cas.NewHashData([]byte(mn.Nonce))
	m["tmstamp"] = cas.DataHasher(cas.TimeHasher(mn.Tmstamp))
	m["txnBlock"] = cas.NewHashData([]byte(MerkleHashFromTxnBlocks(mn.TxnBlock)))
	return m
}

// NewMerkleNode inits a new Block
func NewMerkleNode() *MerkleNode {
	return &MerkleNode{TreeID: GenUUID(), Parent: GenesisMerkleNodeHash(), Nonce: genNonce(), Tmstamp: time.Now(), TxnBlock: make([]BlockTxn, 0)}
}

// MerkleNodeClassName signature namespace
const MerkleNodeClassName = "biz.qzip.bloc.merkle.v01"

// GenesisMerkleNodeHash generates a marker from MerkleBlockClassName
func GenesisMerkleNodeHash() cas.HashData {
	return *cas.NewHashData([]byte(MerkleNodeClassName))
}

// MerkleNodeCutter cuts merkle blocks
type MerkleNodeCutter struct {
	RootHash  cas.Hashable
	BlockSize int
	In        chan *BlockTxn
	Out       chan *MerkleKV
}

// MerkleKV Key is the current Hash & Val is the Block
type MerkleKV struct {
	Key cas.Hashable `json:"key"`
	Val *MerkleNode  `json:"val"`
}

// Process cuts the block
func (bc *MerkleNodeCutter) Process(ctx context.Context) {
	node := NewMerkleNode()
	node.Parent = bc.RootHash
	nodeOp := func() {
		node.Tmstamp = time.Now()
		mv := &MerkleKV{
			Val: node,
			Key: node,
		}
		bc.Out <- mv
		node = NewMerkleNode()
		node.Parent = mv.Key
	}
	for txn := range bc.In {
		//TODO: insert context close logic. via switch
		node.TxnBlock = append(node.TxnBlock, *txn)
		if len(node.TxnBlock) >= bc.BlockSize {
			nodeOp()
		}
	} // for
	if len(node.TxnBlock) > 0 {
		nodeOp()
	}
	close(bc.Out)
}

func genNonce() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		b = []byte("12345678901234567890123456789012")
	}
	return hex.EncodeToString(b)
}

//GenUUID Generate UUID
func GenUUID() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		uuid = []byte("c0ffeebeebeeff0c")

	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid)
}
