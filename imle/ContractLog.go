package imle

import (
	"crypto/sha512"
	//"encoding/json"
	"fmt"
	"merkle"
	"strings"
	"time"

	"github.com/docker/go/canonical/json"
)

// MerkleDataLog Merkle Data of a Log.
// enables for role based selective data element access
// partial data can be verified against merkle proof
type MerkleDataLog struct {
	Meta       []Trie    `json:"meta"`
	MerkleRoot []byte    `json:"merkleRoot"`
	Data       []Trie    `json:"data"`
	Timestamp  time.Time `json:"tmstamp"`
}

// MerkleProof generates the merkle proof of the Data Log
func (md *MerkleDataLog) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	for _, ele := range md.Data {
		m[ele.Key.String()] = ele.Value
	}
	return merkle.SimpleProofsFromMap(m)
}

// Trie TLV (Type Label Value) for merkle Hash,
// key = namespace : label
// value is used to compute the Hash
// namespace is used for JSON serialization
type Trie struct {
	Key   Key    `json:"key"`
	Value Hasher `json:"value"`
}

// TrieSerializer serializes a Trie based on namespace
type TrieSerializer interface {
	DeSerialize(val interface{}) (Hasher, error)
	Serialize(val Hasher) (interface{}, error)
}

var trieRegistry = make(map[string]TrieSerializer, 0)

// RegistryTrieSerializer register the serializer
func RegistryTrieSerializer(ns string, ts TrieSerializer) {
	trieRegistry[ns] = ts
}

// UnmarshalJSON implements https://golang.org/pkg/encoding/json/
func (t *Trie) UnmarshalJSON(b []byte) (err error) {
	des := make(map[string]interface{}, 0)
	err = json.Unmarshal(b, &des)
	if err != nil {
		return
	}
	t.Key, err = t.deserializeKey(des)
	if err != nil {
		return
	}
	v, ok := des["value"]
	if !ok {
		err = fmt.Errorf("Trie.UnmarshalJSON: value not present in %v", des)
		return
	}
	if v == nil {
		err = fmt.Errorf("Trie.UnmarshalJSON: value is empty in %v", des)
		return
	}
	t.Value, err = t.deserializeValue(t.Key, v)
	return
}

func (t *Trie) deserializeKey(des map[string]interface{}) (key Key, err error) {

	k, ok := des["key"]
	if !ok {
		err = fmt.Errorf("Trie.UnmarshalJSON: key not present in %v", des)
		return
	}
	if k == nil {
		err = fmt.Errorf("Trie.UnmarshalJSON: key is empty in %v", des)
		return
	}

	switch v := k.(type) {
	case string:
		if strings.Index(v, ":") > 0 {
			nk := strings.Split(v, ":")
			key.Namespace = nk[0]
			key.Label = strings.Join(nk[1:], "")
		} else {
			key.Namespace = "default"
			key.Label = v
		}

	default:
		if keyval, er := json.MarshalCanonical(k); er == nil {
			err = json.Unmarshal(keyval, &key)
		} else {
			err = er
		}
	}

	return
}

func (t *Trie) deserializeValue(key Key, val interface{}) (h Hasher, err error) {
	ts, ok := trieRegistry[key.Namespace]
	if ok {
		h, err = ts.DeSerialize(val)
	} else {
		h, err = t.convert2hasher(val)
	}
	return
}

func (t *Trie) convert2hasher(v interface{}) (h Hasher, err error) {

	//val, err := json.Marshal(v)
	val, err := json.MarshalCanonical(v)
	if err == nil {
		h = DataHasher(val) //TODO: convert to canonical hash
	}

	return
}

// MarshalJSON implements https://golang.org/pkg/encoding/json/
func (t *Trie) MarshalJSON() (b []byte, err error) {
	tm := make(map[string]interface{})
	tm["key"] = t.Key.String()
	ts, ok := trieRegistry[t.Key.Namespace]
	if !ok {
		tm["value"] = t.Value
	} else {
		tm["value"], err = ts.Serialize(t.Value)
	}
	if err == nil {
		b, err = json.MarshalCanonical(tm) // json.Marshal(tm)
	}
	return
}

// Hasher computes hash of an object
type Hasher interface {
	Hash() []byte
}

// Key namespaced label, the namespace is used for serializing/deserializing
type Key struct {
	Namespace string `json:"namespace"`
	Label     string `json:"label"`
}

func (k Key) String() string {
	return k.Namespace + ":" + k.Label
}

// CasAddress CAS address
type CasAddress []byte

// Hash implements merkle.Hasher
func (h CasAddress) Hash() []byte {
	return ([]byte(h))
}

// DataHasher data that partakes in Merkle Hashing
type DataHasher []byte

// Hash implements merkle.Hasher
func (h DataHasher) Hash() []byte {
	return Hash([]byte(h))
}

// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}
