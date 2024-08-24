package db

import (
	"crypto/sha256"
	"fmt"
	"merkle"
	"time"

	bc "bc/blockchain"
	"bc/cas"
)

/*
		Assumes Table structure as :
	    CREATE TABLE IF NOT EXISTS cas (
			w3cdid	VARCHAR(512) NOT NULL PRIMARY KEY,
			namespace	VARCHAR(512) DEFAULT ('com.aloagri.harvest2invoice'),
			cas_data	JSON,
			tmstamp	TIMESTAMP
		);
*/
type DbCasSchema struct {
	W3cDid    string      `json:"w3cdid"`
	Namespace string      `json:"namespace"`
	CasData   interface{} `json:"cas_data"`
	Tmstamp   time.Time   `json:"tmstamp"`
}

type TxnQueue struct {
	ChainID         string       `json:"chainID"`
	NextBlockHeight int          `json:"nextHeight"`
	Validator       bc.Signature `json:"validator"`
	Txn             []LogHash    `json:"pendingTransactions"`
}

type LogHash []byte

func (l *LogHash) Hash() []byte { return []byte(*l) }

func LogHashArray(logs []LogHash) []byte {
	hash := sha256.New()
	for _, i := range logs {
		hash.Write([]byte(i))
	}
	return hash.Sum(nil)
}

func (tq *TxnQueue) MerkleHasher() map[string]merkle.Hasher {
	m := make(map[string]merkle.Hasher)
	m["chainID"] = cas.NewHashData([]byte(tq.ChainID))
	m["nextHeight"] = cas.NewHashData([]byte(fmt.Sprintf("%d", tq.NextBlockHeight)))
	m["validator"] = cas.NewHashData(tq.Validator.Hash())
	m["pendingTransactions"] = cas.NewHashData(LogHashArray(tq.Txn))
	return m
}

// copied from tm2/pkg/db/types.go
// DBs are goroutine safe.
type DB interface {
	// Get returns nil iff key doesn't exist.
	// A nil key is interpreted as an empty byteslice.
	// CONTRACT: key, value readonly []byte
	Get([]byte) []byte

	// Has checks if a key exists.
	// A nil key is interpreted as an empty byteslice.
	// CONTRACT: key, value readonly []byte
	Has(key []byte) bool

	// Set sets the key.
	// A nil key is interpreted as an empty byteslice.
	// CONTRACT: key, value readonly []byte
	Set([]byte, []byte)
	SetSync([]byte, []byte)

	// Delete deletes the key.
	// A nil key is interpreted as an empty byteslice.
	// CONTRACT: key readonly []byte
	Delete([]byte)
	DeleteSync([]byte)

	// Iterate over a domain of keys in ascending order. End is exclusive.
	// Start must be less than end, or the Iterator is invalid.
	// A nil start is interpreted as an empty byteslice.
	// If end is nil, iterates up to the last item (inclusive).
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// CONTRACT: start, end readonly []byte
	Iterator(start, end []byte) Iterator

	// Iterate over a domain of keys in descending order. End is exclusive.
	// Start must be less than end, or the Iterator is invalid.
	// If start is nil, iterates up to the first/least item (inclusive).
	// If end is nil, iterates from the last/greatest item (inclusive).
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// CONTRACT: start, end readonly []byte
	ReverseIterator(start, end []byte) Iterator

	// Closes the connection.
	Close()

	// Creates a batch for atomic updates.
	NewBatch() Batch

	// For debugging
	Print()

	// Stats returns a map of property values for all keys and the size of the cache.
	Stats() map[string]string
}

//----------------------------------------
// Batch

// Batch Close must be called when the program no longer needs the object.
type Batch interface {
	SetDeleter
	Write()
	WriteSync()
	Close()
}

type SetDeleter interface {
	Set(key, value []byte) // CONTRACT: key, value readonly []byte
	Delete(key []byte)     // CONTRACT: key readonly []byte
}

//----------------------------------------
// Iterator

/*
Usage:

var itr Iterator = ...
defer itr.Close()

	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(); itr.Value()
		// ...
	}
*/
type Iterator interface {
	// The start & end (exclusive) limits to iterate over.
	// If end < start, then the Iterator goes in reverse order.
	//
	// A domain of ([]byte{12, 13}, []byte{12, 14}) will iterate
	// over anything with the prefix []byte{12, 13}.
	//
	// The smallest key is the empty byte array []byte{} - see BeginningKey().
	// The largest key is the nil byte array []byte(nil) - see EndingKey().
	// CONTRACT: start, end readonly []byte
	Domain() (start []byte, end []byte)

	// Valid returns whether the current position is valid.
	// Once invalid, an Iterator is forever invalid.
	Valid() bool

	// Next moves the iterator to the next sequential key in the database, as
	// defined by order of iteration.
	//
	// If Valid returns false, this method will panic.
	Next()

	// Key returns the key of the cursor.
	// If Valid returns false, this method will panic.
	// CONTRACT: key readonly []byte
	Key() (key []byte)

	// Value returns the value of the cursor.
	// If Valid returns false, this method will panic.
	// CONTRACT: value readonly []byte
	Value() (value []byte)

	// Close releases the Iterator.
	Close()
}
