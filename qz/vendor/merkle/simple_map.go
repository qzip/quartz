package merkle

//	"github.com/cometbft/cometbft/crypto/tmhash"

// Merkle tree from a map.
// Leaves are `hash(key) | hash(value)`.
// Leaves are sorted before Merkle hashing.
type SimpleMap struct {
	kvs    KVPairs
	sorted bool
}

func NewSimpleMap() *SimpleMap {
	return &SimpleMap{
		kvs:    nil,
		sorted: false,
	}
}

// Set hashes the key and value and appends it to the kv pairs.
func (sm *SimpleMap) Set(key string, value Hasher) {
	sm.sorted = false

	// The value is hashed, so you can
	// check for equality with a cached value (say)
	// and make a determination to fetch or not.
	vhash := value.Hash()

	sm.kvs = append(sm.kvs, KVPair{
		Key:   []byte(key),
		Value: vhash,
	})
}

// Hash Merkle root hash of items sorted by key
// (UNSTABLE: and by value too if duplicate key).
func (sm *SimpleMap) Hash() []byte {
	sm.Sort()
	return hashKVPairs(sm.kvs)
}

func (sm *SimpleMap) Sort() {
	if sm.sorted {
		return
	}
	sm.kvs.Sort()
	sm.sorted = true
}

// Returns a copy of sorted KVPairs.
// NOTE these contain the hashed key and value.
func (sm *SimpleMap) KVPairs() KVPairs {
	sm.Sort()
	kvs := make(KVPairs, len(sm.kvs))
	copy(kvs, sm.kvs)
	return kvs
}

//----------------------------------------
/*
// A local extension to KVPair that can be hashed.
// Key and value are length prefixed and concatenated,
// then hashed.

type KVPair cmn.KVPair

func (kv KVPair) Hash() []byte {
	hasher := tmhash.New()
	err := encodeByteSlice(hasher, kv.Key)
	if err != nil {
		panic(err)
	}
	err = encodeByteSlice(hasher, kv.Value)
	if err != nil {
		panic(err)
	}
	return hasher.Sum(nil)
}

func hashKVPairs(kvs KVPairs) []byte {
	kvsH := make([]Hasher, len(kvs))
	for i, kvp := range kvs {
		kvsH[i] = KVPair(kvp)
	}
	return SimpleHashFromHashers(kvsH)
}
*/
