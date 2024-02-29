package cas

import (
	//"crypto/sha256"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"
	//"github.com/tendermint/go-crypto"
)

type Hashable interface {
	Hash() []byte
}

//TODO: cleanup reduntant objects

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

// Hash returns SHA-2 SHA-256 hash
func Hash(in []byte) []byte {
	hash := sha256.New()
	hash.Write(in)
	return hash.Sum(nil)
}

func TimeHasher(t time.Time) []byte {
	hash := sha256.New()
	b, _ := t.MarshalBinary()
	hash.Write(b)
	return hash.Sum(nil)
}
func IntHasher(i int64) []byte {
	hash := sha256.New()
	hash.Write(Int2Bin(i))
	return hash.Sum(nil)
}
func IntArrHasher(iarr []int64) []byte {
	hash := sha256.New()
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
	hash := sha256.New()
	hash.Write(buf.Bytes())
	return hash.Sum(nil)
}
func BoolHasher(i bool) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	hash := sha256.New()
	hash.Write(buf.Bytes())
	return hash.Sum(nil)
}

//------------------- old qz code

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

/*
// Hash returns SHA-256 hash (tendermint merkle tree uses SHA256)
func Hash(in []byte) []byte {
	hash := sha256.New()
	hash.Write(in)
	return hash.Sum(nil)
}

  merkle hasher uses SHA256
// Hash returns SHA-2 SHA-256 hash
func Hash(in []byte) []byte {
	hash := sha256.New()
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

// Hash returns SHA-2 SHA-256 hash
func Hash(in []byte) []byte {
	hash := sha256.New()
	hash.Write(in)
	return hash.Sum(nil)
}
*/
