package blockchain

/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

import (
	"context"
	"crypto/sha256"

	//"crypto/sha512"
	"bc/cas"
	"fmt"
	"merkle"
	"time"
)

/*
Blockchain is a special case of txn chain. At a regular interval the blockchain process wakes up,
selects the transactions that occured between now and the last block; then creates a record:

Blocks are created and stored as txnchain (blockchain)

Signing is detached and save as seperate records. Thus a block can
have multiple signers

*/

// https://github.com/alecthomas/jsonschema
// jsonschema="title=the name,description=The name of a friend,example=joe,example=lucy,default=alex"`

// SignedBlockChain in memory respresenation of Blockchain record
type SignedBlockChain struct {
	Head  BlockChain  `json:"chainHead"`
	Signs []Signature `json:"signs,omitempty"`
}

// Signature JWS of hashed data
type Signature struct {
	ContentHash []byte `json:"contentHash"` // for chain head it's chainhead hash
	Jws         []byte `json:"jws"`
	JwkPub      []byte `json:"jwk"`
}

// SignedData Data + signature
type SignedData struct {
	Sig      Signature  `json:"signature"`
	Content  SigContent `json:"signedContent"`
	PermRole string     `json:"permRole"`
	PartData []byte     `json:"partData,omitempty"` // SignContext provides info on Json schema
}

// SigContent content of the signature
type SigContent struct {
	DataURI  string    `json:"dataURI"`
	RefID    string    `json:"refID"`
	FullHash []byte    `json:"fullHash"`
	PartHash []byte    `json:"partHash,omitempty"`
	Tmstamp  time.Time `json:"tmstamp"`
}

// BlockChain memory representation of block that is signed and placed in Blockchain record
type BlockChain struct {
	ChainID       string     `json:"chainID"`
	Height        int64      `json:"height"`
	PrevChainHash []byte     `json:"prevChainHash"` // genesis block previous hash is Hash(ChainID)
	Nonce         string     `json:"nonce"`
	Tmstamp       time.Time  `json:"tmstamp"`
	TxnBlock      []BlockTxn `json:"txnBlock"`
}

// Hash implements merkle.Hasher interface
// computes merkle root of the Block
func (blk *BlockChain) Hash() []byte {
	return merkle.SimpleHashFromMap(blk.MerkleHasher())
}

// MerkleProofs computes Hash & proof of array of TxnBlock(s)
func (blk *BlockChain) MerkleProofs() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	return merkle.SimpleProofsFromMap(blk.MerkleHasher())
}

// MerkleHasher transforms BlockTxn into merkle hasher map
func (blk *BlockChain) MerkleHasher() map[string]merkle.Hasher {
	m := make(map[string]merkle.Hasher)
	m["chainID"] = cas.NewHashData([]byte(blk.ChainID))
	m["height"] = cas.NewHashData([]byte(fmt.Sprintf("%d", blk.Height)))
	m["prevChainHash"] = cas.NewHashData(blk.PrevChainHash)
	m["nonce"] = cas.NewHashData([]byte(blk.Nonce))
	m["tmstamp"] = cas.NewHashData([]byte(blk.Tmstamp.String()))
	m["txnBlock"] = cas.NewHashData([]byte(MerkleHashFromTxnBlocks(blk.TxnBlock)))
	return m
}

// BlockTxn txn meta data part of the block
type BlockTxn struct {
	TxnRef     string       `json:"txnRef"`
	TxnHash    cas.Hashable `json:"txnHash"`
	TxnFamily  cas.Hashable `json:"txnFamilyHash"` // Hash of ClassName
	TxnTmstamp time.Time    `json:"txnTmstamp"`
}

// Hash implements merkle.Hasher interface
// computes merkle root of the BxnBlock param
func (btx *BlockTxn) Hash() []byte {
	return merkle.SimpleHashFromMap(btx.MerkleHasher())

}

// MerkleHasher transforms BlockTxn into merkle hasher map
func (btx *BlockTxn) MerkleHasher() map[string]merkle.Hasher {
	m := make(map[string]merkle.Hasher)
	m["txnRef"] = cas.NewHashData([]byte(btx.TxnRef))
	m["txnHash"] = btx.TxnHash
	m["txnFamilyHash"] = btx.TxnFamily
	m["txnTmstamp"] = cas.DataHasher(cas.TimeHasher(btx.TxnTmstamp))
	return m
}

// MerkleProofsFromTxnBlocks computes Hash & proof of array of TxnBlock(s)
func MerkleProofsFromTxnBlocks(txnBlock []BlockTxn) (rootHash []byte, proofs []*merkle.SimpleProof) {
	mhsh := make([]merkle.Hasher, len(txnBlock))
	for i, t := range txnBlock {
		mhsh[i] = &t
	}
	return merkle.SimpleProofsFromHashers(mhsh)
}

// MerkleHashFromTxnBlocks computes Hash of array of TxnBlock(s)
func MerkleHashFromTxnBlocks(txnBlock []BlockTxn) []byte {
	mhsh := make([]merkle.Hasher, len(txnBlock))
	for i, t := range txnBlock {
		mhsh[i] = &t
	}
	return merkle.SimpleHashFromHashers(mhsh)
}

// JSONCas Content Addressable Storage
type JSONCas struct {
	Hash      []byte                 `json:"hash"`
	ClassName string                 `json:"className"`
	Data      map[string]interface{} `json:"data"`
	Tmstamp   time.Time              `json:"tmstamp"`
}

// NewBloc inits a new Block
func NewBloc() *BlockChain {
	return &BlockChain{Nonce: cas.GenNonce(), Tmstamp: time.Now(), TxnBlock: make([]BlockTxn, 0)}
}

// ChainHash computes HASH  SHA-256 (downgraded from SHA-512) as tendermint merkle uses SHA256
// ChainHash computes HASH  SHA-2 512 (64 bytes hash) 128 hex string returned
func ChainHash(cur, last []byte) (hash []byte) {

	//hasher := sha512.New()
	hasher := sha256.New()
	hasher.Write(last)
	hasher.Write(cur)
	hash = hasher.Sum(nil)

	return
}

// BlockCutter cuts the block
type BlockCutter struct {
	Helper BlockHelper
	In     chan *BlockTxn
	Out    chan *BlockChain
	Err    chan error
}

// BlockHelper block height keeper
type BlockHelper interface {
	ComputeNextHeight(ctx context.Context, echan chan error) int64
	LastChainHash() cas.Hashable
	ChainID() string
	BlockSize() int
}

// Process cuts the block
func (bc *BlockCutter) Process(ctx context.Context) {
	blksiz := bc.Helper.BlockSize()
	//txnBlock := make([]BlockTxn,0)
	blk := NewBloc()
	for txn := range bc.In {
		blk.TxnBlock = append(blk.TxnBlock, *txn)
		if len(blk.TxnBlock) >= blksiz {
			blk.Height = bc.Helper.ComputeNextHeight(ctx, bc.Err)
			blk.PrevChainHash = bc.Helper.LastChainHash().Hash()
			blk.ChainID = bc.Helper.ChainID()
			blk.Tmstamp = time.Now()
			bc.Out <- blk
			blk = NewBloc()
		}
	}
	if len(blk.TxnBlock) > 0 {
		blk.Height = bc.Helper.ComputeNextHeight(ctx, bc.Err)
		blk.PrevChainHash = bc.Helper.LastChainHash().Hash()
		blk.ChainID = bc.Helper.ChainID()
		blk.Tmstamp = time.Now()
		bc.Out <- blk
	}
	close(bc.Out)
}
