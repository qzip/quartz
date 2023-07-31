package cas

import (
	"encoding/json"
	"fmt"
	"time"

	//"golang.org/x/crypto/sha3"
	canjson "github.com/docker/go/canonical/json"
)

/*
// Cas interface used for abstracting the database storage
type Cas interface {
	PutJson(ctx context.Context, data *CasJSON) (err error)
}
*/

// CasJSON Content Addressable Storage
type CasJSON struct {
	Hash      HashData               `json:"hash"`
	Namespace string                 `json:"namespace"`
	Key       string                 `json:"key"`
	Data      map[string]interface{} `json:"data"`
	Tmstamp   time.Time              `json:"tmstamp"`
}

// SignedCasJSON detached signed CAS
type SignedCasJSON struct {
	Jws string  `json:"jws"` // JWS with embedded hash & JWK
	Cas CasJSON `json:"sig"`
}

// MapToSignedCas  JSON structure to signed CAS transformer
type MapToSignedCas struct {
	Namespace   string
	In          chan map[string]interface{}
	Out         chan *SignedCasJSON
	Err         chan error
	KeyName     string
	JwkFileName string // if not set embedded JWK is used
}

// Process transforms JSON structure to signed CAS
func (c *MapToSignedCas) Process() {
	for in := range c.In {
		out := &SignedCasJSON{

			Cas: CasJSON{
				Hash:      *NewHashData(c.ComputeHash(in)),
				Namespace: c.Namespace,
				Key:       fmt.Sprint(in[c.KeyName]),
				Data:      in,
				Tmstamp:   time.Now(),
			},
		}
		out.Jws = c.Sign(out.Cas.Hash)
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
