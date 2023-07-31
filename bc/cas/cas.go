package cas

import (
	"merkle"
	"time"
	//"golang.org/x/crypto/sha3"
)

/*
// Cas interface used for abstracting the database storage
type Cas interface {
	PutJson(ctx context.Context, data *CasJSON) (err error)
}
*/

// CasJSON Content Addressable Storage
type CasJSON struct {
	Cas     Cas       `json:"cas"`
	Key     string    `json:"key"`
	Tmstamp time.Time `json:"tmstamp"`
}

//Hash implements merkle.Hasher interface
// computes merkle root of the Block
func (cj *CasJSON) Hash() []byte {
	return merkle.SimpleHashFromMap(cj.MerkleHasher())
}

//MerkleProofs computes Hash & proof of array of TxnBlock(s)
func (cj *CasJSON) MerkleProofs() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	return merkle.SimpleProofsFromMap(cj.MerkleHasher())
}

//MerkleHasher transforms BlockTxn into merkle hasher map
func (cj *CasJSON) MerkleHasher() map[string]merkle.Hasher {
	m := make(map[string]merkle.Hasher)
	m["cas"] = cj.Cas
	m["key"] = NewHashData([]byte(cj.Key))
	m["tmstamp"] = HashedData(TimeHasher(cj.Tmstamp))
	return m
}

//Cas Content addressable store abstraction
type Cas interface {
	Hash() []byte
	UnmarshalJSON(b []byte) error
	MarshalJSON() ([]byte, error)
	Namespace() string
}

var registry = make(map[string]Cas)

// RegisterCas commands must register to be visible to the commander
func RegisterCas(cmd Cas) {
	registry[cmd.Namespace()] = cmd
}

// LookUpCas gets command from registry
func LookUpCas(cmd string) Cas {
	return registry[cmd]
}

// IterateCas for each registered command call back
func IterateCas(callBack func(Cas)) {
	for _, v := range registry {
		callBack(v)
	}
}

// SignedCasJSON detached signed CAS
type SignedCasJSON struct {
	Jws string  `json:"jws"` // JWS with embedded hash & JWK
	Cas CasJSON `json:"sig"`
}

// MapToSignedCas  JSON structure to signed CAS transformer
type MapToSignedCas struct {
	Namespace   string
	In          chan *CasJSON
	Out         chan *SignedCasJSON
	Err         chan error
	KeyName     string
	JwkFileName string // if not set embedded JWK is used
}

// Process transforms JSON structure to signed CAS
func (c *MapToSignedCas) Process() {
	for in := range c.In {
		out := &SignedCasJSON{
			Cas: *in,
		}
		out.Jws = c.Sign(out.Cas.Hash())
		c.Out <- out
	}
	close(c.Out)

}

// Sign the hash
func (c *MapToSignedCas) Sign(hash HashData) string {
	jwk, err := GetJSONWebKey(c.JwkFileName)
	if err != nil {
		c.Err <- err // TODO: should err chan be closed?
		return ""
	}
	sig, err := SignHash(jwk, []byte(hash))
	if err != nil {
		c.Err <- err // TODO: should err chan be closed?
		return ""
	}
	return sig
}

/*
// ComputeHash tries to compute canonical hash first, if failed, the
// normal json marshalled hash is computed.
func (c *MapToSignedCas) ComputeHash(inJSON map[string]interface{}) (out []byte) {
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
*/
