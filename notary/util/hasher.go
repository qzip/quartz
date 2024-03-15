package util

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"time"
)

//HashedData pass through wrapper
type HashedData []byte

//Hash implements merkle.Hasher
func (h HashedData) Hash() []byte {
	return h
}

//DataHasher data that partakes in Merkle Hashing
type DataHasher []byte

//Hash implements merkle.Hasher
func (h DataHasher) Hash() []byte {
	return Hash([]byte(h))
}

// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}

func TimeHasher(t time.Time) []byte {
	hash := sha512.New()
	b, _ := t.MarshalBinary()
	hash.Write(b)
	return hash.Sum(nil)
}
func IntHasher(i int64) []byte {
	hash := sha512.New()
	hash.Write(Int2Bin(i))
	return hash.Sum(nil)
}
func IntArrHasher(iarr []int64) []byte {
	hash := sha512.New()
	for _, i := range iarr {
		hash.Write(Int2Bin(i))
	}
	return hash.Sum(nil)
}
func Int2Bin(i int64) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func Uint8Hasher(i uint8) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	hash := sha512.New()
	hash.Write(buf.Bytes())
	return hash.Sum(nil)
}
func BoolHasher(i bool) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	hash := sha512.New()
	hash.Write(buf.Bytes())
	return hash.Sum(nil)
}
