package util

import (
	"bc/cas"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"

	canjson "github.com/docker/go/canonical/json"
	jose "gopkg.in/square/go-jose.v2"
)

/*
gen key : src/gopkg.in/square/go-jose.v2/jwk-keygen/main.go


 jwk-keygen --use=sig --alg=EdDSA --kid="in.qzip.blockchain.01may2019"


*/

var jwkBin = []byte(`{"use":"sig","kty":"OKP","kid":"in.qzip.blockchain.01may2019","crv":"Ed25519","alg":"EdDSA","x":"e2vtHFZafjEDB_CH4R0NRSFke5M0Ls0DsPUDy_na4yo","d":"rz-k_Sli54v4yv9pMCFTnzfCaPBXlOjwgSr6A_VDHf8"}`)

// CanonicalBinJSON input json content, output Canonical json bin
// "Floating point numbers are not allowed in canonical JSON. Neither are leading zeros or "minus 0" for integers."
// Read more at: http://wiki.laptop.org/go/Canonical_JSON
func CanonicalBinJSON(in []byte) (out []byte, err error) {
	inJSON := make(map[string]interface{}, 0)
	err = json.Unmarshal(in, &inJSON)
	if err != nil {
		return
	}
	out, err = canjson.MarshalCanonical(inJSON)

	return
}

/*
// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}
*/

// CanonicalHash JSON is canonilized and then hashed
// TODO: Canonical JSON gives error for float amount.
// as a fix, removed cannonical marshalling in case of error
func CanonicalHash(in []byte) (out []byte, err error) {
	ch, er := CanonicalBinJSON(in)
	if er != nil {
		//err = er
		ch = in
	}

	out = cas.Hash(ch)

	return
}

// Bin2Str converts byte array to Hex string
func Bin2Str(in []byte) (string, error) {
	return hex.EncodeToString(in), nil
}

// Str2Bin Hex string to binary
func Str2Bin(in string) (out []byte, err error) {
	out, err = hex.DecodeString(in)
	return
}

/*
// Bin2Str converts byte array to JSON string
func Bin2Str(in []byte) (out string, err error) {
	b, err := json.Marshal(in)
	if err == nil {
		out = string(b)
	}
	return
}

// Str2Bin JSON string to binary
func Str2Bin(in string) (out []byte, err error) {
	out = make([]byte, 0)
	err = json.Unmarshal([]byte(in), &out)
	return
}
*/

// GetJSONWebKey gets a JSONWebKey from file
func GetJSONWebKey(jwkFile string) (jwk *jose.JSONWebKey, err error) {
	var jwkIn = jwkBin
	if len(jwkFile) > 1 {
		jwkIn, err = ioutil.ReadFile(jwkFile)
		if err != nil {
			return
		}
	}
	jwk = &jose.JSONWebKey{}
	err = jwk.UnmarshalJSON([]byte(jwkIn))
	return
}

// SignHash returns fully Serialized JSONWebSignature with embedded JSONWebKey (Public)
func SignHash(jwk *jose.JSONWebKey, hash []byte) (sig string, err error) {
	ngen := &Noncer{}
	so := &jose.SignerOptions{NonceSource: ngen, EmbedJWK: true}
	so.WithContentType("SHA512")
	sk := jose.SigningKey{Algorithm: jose.EdDSA, Key: jwk}
	si, err := jose.NewSigner(sk, so)
	if err != nil {
		return
	}

	jws, err := si.Sign(hash)
	if err != nil {
		return
	}

	sig = jws.FullSerialize()
	return
}

// VerifySig Verifies the JWS with the passed key. If key is nil, then uses embedded sig (UNSAFE, NOT RECOMMENDED)
func VerifySig(input string, jwk *jose.JSONWebKey) (payload []byte, err error) {
	jws, err := jose.ParseSigned(input)
	if err != nil {
		return
	}

	if jwk != nil {
		pub := jwk.Public()
		payload, err = jws.Verify(&pub)
	} else {
		// unsafe, as anyone might have signed it.
		payload, err = jws.Verify(jws.Signatures[0])
	}

	return
}

// Noncer nonce generator
type Noncer struct {
}

// Nonce creates a new instance of Noncer
func (*Noncer) Nonce() (nonce string, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)

	nonce = hex.EncodeToString(b)
	return
}
