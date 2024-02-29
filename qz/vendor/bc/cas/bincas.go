package cas

import (
	"time"
)

// BinCas Content Addressable Storage
type BinCas struct {
	Hash      HashData  `json:"hash"`
	Namespace string    `json:"namespace"`
	Key       string    `json:"key"`
	Data      []byte    `json:"data"`
	Tmstamp   time.Time `json:"tmstamp"`
	Mimetype  string    `json:"mimetype"`
}

// SignedBinCas detached signed CAS
type SignedBinCas struct {
	Jws string `json:"jws"` // JWS with embedded hash & JWK
	Cas BinCas `json:"sig"`
}

// MapToSignedBinCas  JSON structure to signed CAS transformer
type MapToSignedBinCas struct {
	Namespace   string
	In          chan *BinCasKeyVal
	Out         chan *SignedBinCas
	Err         chan error
	JwkFileName string // if not set embedded JWK is used
}

// BinCasKeyVal channel data
type BinCasKeyVal struct {
	Key      string
	Val      []byte
	Mimetype string
}

// Process transforms JSON structure to signed CAS
func (c *MapToSignedBinCas) Process() {
	for in := range c.In {
		out := &SignedBinCas{

			Cas: BinCas{
				Hash:      *NewHashData(Hash(in.Val)),
				Namespace: c.Namespace,
				Key:       in.Key,
				Data:      in.Val,
				Mimetype:  in.Mimetype,
				Tmstamp:   time.Now(),
			},
		}
		out.Jws = c.Sign(out.Cas.Hash)
		c.Out <- out
	}
	close(c.Out)

}

// Sign the hash
func (c *MapToSignedBinCas) Sign(hash HashData) string {
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

// ComputeHash computes hash.
func (c *MapToSignedBinCas) ComputeHash(b []byte) (out []byte) {
	out = Hash(b)
	return
}
