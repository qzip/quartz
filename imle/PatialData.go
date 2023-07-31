package imle

import "merkle"

//PartialData partial txn data, for state transaction observers
// with limited visibility
type PartialData struct {
	ParialData     *MerkleDataLog
	MaskedElements map[string]CasAddress
}

//NewPartialData instantiates Partial Data object
func NewPartialData(data *MerkleDataLog, maskFields []string) *PartialData {
	dl := &MerkleDataLog{
		Meta:       data.Meta,
		MerkleRoot: data.MerkleRoot,
		Data:       make([]Trie, 0),
		Timestamp:  data.Timestamp,
	}
	pd := &PartialData{
		ParialData:     dl,
		MaskedElements: make(map[string]CasAddress),
	}
	zero := CasAddress(make([]byte, 0))
	for _, mf := range maskFields {
		pd.MaskedElements[mf] = zero
	}
	for _, ele := range dl.Data {
		_, ok := pd.MaskedElements[ele.Key.String()]
		if ok {
			// masked filed
			pd.MaskedElements[ele.Key.String()] = ele.Value.Hash()
		} else {
			pd.ParialData.Data = append(pd.ParialData.Data, ele)
		}
	}
	return pd
}

//MerkleProof generates the merkle proof of the Data Log
func (md *PartialData) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	for _, ele := range md.ParialData.Data {
		m[ele.Key.String()] = ele.Value
	}
	for k, v := range md.MaskedElements {
		m[k] = v
	}
	return merkle.SimpleProofsFromMap(m)
}
