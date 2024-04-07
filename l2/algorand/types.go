package algorand

import (
	"bc/cas"
	"merkle"
	"time"
)

type AlgoCommitment struct {
	Schema     string    `json:"schema"`
	Namespace  string    `json:"namespace"`
	State      string    `json:"state,omitempty"`
	Sequence   int64     `json:"sequence,omitempty"`
	Nonce      string    `json:"nonce,omitempty"`
	Tmstamp    time.Time `json:"tmstamp"`
	MerkleRoot []byte    `json:"merkle"`
}

// Hash implements merkle.Hasher interface
// computes merkle root of the TxnBlock param
func (ac *AlgoCommitment) Hash() []byte {
	return merkle.SimpleHashFromMap(ac.MerkleHasher())

}

func (ac *AlgoCommitment) MerkleHasher() map[string]merkle.Hasher {
	m := make(map[string]merkle.Hasher)

	m["schema"] = cas.HashedData(cas.Hash([]byte(ac.Schema)))
	m["namespace"] = cas.HashedData(cas.Hash([]byte(ac.Namespace)))
	if ac.Sequence > 0 {
		m["sequence"] = cas.HashedData(cas.IntHasher(ac.Sequence))
	}
	if len(ac.State) > 0 {
		m["state"] = cas.HashedData(cas.Hash([]byte(ac.State)))
	}
	if len(ac.Nonce) > 0 {
		m["nonce"] = cas.HashedData(cas.Hash([]byte(ac.Nonce)))
	}
	m["tmstamp"] = cas.HashedData(cas.TimeHasher(ac.Tmstamp))
	m["merkle"] = cas.HashedData(ac.MerkleRoot)
	return m

}
