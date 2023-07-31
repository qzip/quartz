package cas

import (
	//"crypto/sha512"
	"crypto/sha256"
	"encoding/hex"
	//"github.com/tendermint/go-crypto"
)

// HashData encapsulates byte array for serializing to JSON
type HashData []byte

// MarshalText custom marshall
func (h *HashData) MarshalText() ([]byte, error) {

	return []byte(hex.EncodeToString(*h)), nil
}
func (h *HashData) String() string {
	return hex.EncodeToString(*h)
}

//Hash implements merkle.Hasher
func (h HashData) Hash() []byte {
	return h
	//Hash([]byte(h))
}

// UnmarshalText custom unmarshall
func (h *HashData) UnmarshalText(in []byte) (err error) {
	b, err := hex.DecodeString(string(in))
	hd := HashData(b)
	*h = hd // TODO: test if the value changes
	// TODO: need to fix the self mutation design

	return
}

// NewHashData encapsulate byte array
func NewHashData(in []byte) *HashData {
	hx := HashData(in)
	return &hx
}

// Hash returns SHA-256 hash (tendermint merkle tree uses SHA256)
func Hash(in []byte) []byte {
	hash := sha256.New()
	hash.Write(in)
	return hash.Sum(nil)
}

/*  merkle hasher uses SHA256
// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}
*/

/*
// NewHashData encapsulate byte array
func NewHashData(in []byte) HashData {
	hx := HashData(Hash(in))
	return hx
}

// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}
*/
