package digicon

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"imle"
	"time"

	canjson "github.com/docker/go/canonical/json"
)

// DigitalContractBlock blocks are chained together.
// each of the blocks are storedin a CAS.
// The Key: $ID and the latest block CAS Address is stored
// in a table. Thus each Digital Contract state changed are
// inserted in a blockchain.
// The entities of the contracts are lumpted in a Map structure
// and Merkelized. The advantage of such an approach is that
// All the elements of the records need not be exposed to all
// the stake holders. For example, each of the record may have
// some confidential information and some public information.
// Though the public can see only the parts of the record yet,
// they will be able to verify the data integrigity by requesting
// a Merkle Proof of the record.
type DigitalContractBlock struct {
	Block     *imle.Block `json:"block"`
	ID        string      `json:"id"`
	Timestamp time.Time   `json:"tmstamp"`
	BlockSig  *Signature  `json:"sig"`
}

// Signature JWS of hashed data
type Signature struct {
	ContentHash HashData               `json:"contentHash"` // for chain head it's chainhead hash
	Jws         map[string]interface{} `json:"jws"`
	JwkPub      map[string]interface{} `json:"jwk"`
}

// DigitalContractBlockStore storage interface for Contract Chain
type DigitalContractBlockStore interface {
	LoadCurrentRoot(contractID string) (*DigitalContractBlock, error)
	SaveCurrentRoot(block *DigitalContractBlock) (cas imle.CasAddress, err error)
}

// ComputeHash computes the HASH of the JSON
func ComputeHash(inJSON map[string]interface{}) (out []byte) {
	b, err := canjson.MarshalCanonical(inJSON)
	if err != nil {
		b, err = json.Marshal(inJSON)
		if err != nil {
			b = []byte(fmt.Sprintf("%v", inJSON))
		}
	}
	out = Hash(b)
	return
}

// HasherJSON implements Hasher
type HasherJSON map[string]interface{}

// Hash implements Hasher
func (h *HasherJSON) Hash(in []byte) []byte {
	return ComputeHash(map[string]interface{}(*h))
}

// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}

// HashData encapsulates byte array for serializing to JSON
type HashData []byte

// MarshalText custom marshall
func (h HashData) MarshalText() ([]byte, error) {

	return []byte(hex.EncodeToString(h)), nil
}
func (h HashData) String() string {
	return hex.EncodeToString(h)
}

// Hash implements merkle.Hasher
func (h HashData) Hash() []byte {
	return Hash([]byte(h))
}

// UnmarshalText custom unmarshall
func (h HashData) UnmarshalText(in []byte) (err error) {
	b, err := hex.DecodeString(string(in))
	h = HashData(b)

	return
}

// NewHashData encapsulate byte array
func NewHashData(in []byte) HashData {
	hx := HashData(in)
	return hx
}
