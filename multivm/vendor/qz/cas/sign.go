package cas

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"

	jose "gopkg.in/square/go-jose.v2"
)

/*
gen key : src/gopkg.in/square/go-jose.v2/jwk-keygen/main.go


 jwk-keygen --use=sig --alg=EdDSA --kid="in.qzip.blockchain.01may2019"


*/

var jwkBin = []byte(`{"use":"sig","kty":"OKP","kid":"in.qzip.blockchain.01may2019","crv":"Ed25519","alg":"EdDSA","x":"e2vtHFZafjEDB_CH4R0NRSFke5M0Ls0DsPUDy_na4yo","d":"rz-k_Sli54v4yv9pMCFTnzfCaPBXlOjwgSr6A_VDHf8"}`)

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
	//so.WithContentType("SHA512")
	so.WithContentType("SHA256") // tendermint uses SHA256
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

// GenNonce generates a nonce string
func GenNonce() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "c0ffeebeebeeff0c"
	}
	return hex.EncodeToString(b)
}

// Noncer nonce generator
type Noncer struct {
}

// Nonce creates a new instance of Noncer
func (*Noncer) Nonce() (nonce string, _ error) {
	nonce = GenNonce()
	return
}
