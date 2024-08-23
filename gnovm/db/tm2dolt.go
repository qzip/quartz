package db

const (
	defaultCommitMsg    = "traceprod immutable ledger"
	defaultCommitAuthor = "Quartz <tech@innomon.in>"
	defaultCommitBranch = "main"
)

// Tm2Dolt implements Tendermint2 Database, implements DB tm2/pkg/db/types.go
// modelled on `tm2/pkg/db/memdb/mem_db.go`
// stores in DoltDb `CAS` table
type Tm2Dolt struct {
}

func (td *Tm2Dolt) Get(key []byte) []byte {

	return nil
}

func (td *Tm2Dolt) Has(key []byte) bool {

	return false
}

// Set sets the key.
// A nil key is interpreted as an empty byteslice.
func (td *Tm2Dolt) Set([]byte, []byte) {

}

func (td *Tm2Dolt) SetSync([]byte, []byte) {

}

// Delete deletes the key.
// A nil key is interpreted as an empty byteslice.
func (td *Tm2Dolt) Delete([]byte)     {}
func (td *Tm2Dolt) DeleteSync([]byte) {}

// Iterate over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// A nil start is interpreted as an empty byteslice.
// If end is nil, iterates up to the last item (inclusive).
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// CONTRACT: start, end readonly []byte
func (td *Tm2Dolt) Iterator(start, end []byte) Iterator {
	return nil // TODO:
}

// Iterate over a domain of keys in descending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// If start is nil, iterates up to the first/least item (inclusive).
// If end is nil, iterates from the last/greatest item (inclusive).
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// CONTRACT: start, end readonly []byte
func (td *Tm2Dolt) ReverseIterator(start, end []byte) Iterator {
	return nil // TODO:
}

// Closes the connection.
func (td *Tm2Dolt) Close() {}

// Creates a batch for atomic updates.
func (td *Tm2Dolt) NewBatch() Batch {
	return nil // TODO:
}

// For debugging
func (td *Tm2Dolt) Print() {}

// Stats returns a map of property values for all keys and the size of the cache.
func (td *Tm2Dolt) Stats() map[string]string {
	return nil // TODO:
}
