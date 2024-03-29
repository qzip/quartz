package merkle

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
)

// Tree is a Merkle tree interface.
type Tree interface {
	Size() (size int)
	Height() (height int8)
	Has(key []byte) (has bool)
	Proof(key []byte) (value []byte, proof []byte, exists bool) // TODO make it return an index
	Get(key []byte) (index int, value []byte, exists bool)
	GetByIndex(index int) (key []byte, value []byte)
	Set(key []byte, value []byte) (updated bool)
	Remove(key []byte) (value []byte, removed bool)
	HashWithCount() (hash []byte, count int)
	Hash() (hash []byte)
	Save() (hash []byte)
	Load(hash []byte)
	Copy() Tree
	Iterate(func(key []byte, value []byte) (stop bool)) (stopped bool)
	IterateRange(start []byte, end []byte, ascending bool, fx func(key []byte, value []byte) (stop bool)) (stopped bool)
}

// Hasher represents a hashable piece of data which can be hashed in the Tree.
type Hasher interface {
	Hash() []byte
}

// StringHasher implements Hasher
type StringHasher string

func (h StringHasher) Hash() []byte {
	hs := sha256.Sum256([]byte(h))
	return hs[:]
}

// ByteHasher implements Hasher
type ByteHasher []byte

func (h ByteHasher) Hash() []byte {
	hs := sha256.Sum256(h)
	return hs[:]
}

//-----------------------------------------------------------------------
//https://github.com/tendermint/tendermint/blob/master/crypto/merkle/types.go

// Uvarint length prefixed byteslice
func encodeByteSlice(w io.Writer, bz []byte) (err error) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], uint64(len(bz)))
	_, err = w.Write(buf[0:n])
	if err != nil {
		return
	}
	_, err = w.Write(bz)
	return
}
