package imle

import (
	"time"
)

//BHash simple []byte wrapper to implement HashableData & merkle.Hasher
type BHash []byte

//Hash implements merkle.Hasher & HashableData method
func (h BHash) Hash() []byte {
	return h
}

//Value implements HashableData method
func (h BHash) Value() []byte {
	return h
}

//HashableData must be implemnted for Caser storage
type HashableData interface {
	Hash() []byte
	Value() []byte
}

//Caser CAS Storage Provider
type Caser interface {
	Store(data HashableData, ns string) (cas CasAddress, tmstamp time.Time, err error)
	Load(cas CasAddress) (data []byte, ns string, tmstamp time.Time, err error)
}

//NamedCasStore much like IPNS or DNS, maps Name to a CAS Address
type NamedCasStore interface {
	HasCas(id string) (bool, error)
	GetCas(id string) (cas CasAddress, tmstamp time.Time, err error)
	PutCas(id string, cas CasAddress) (tmstamp time.Time, err error)
}
