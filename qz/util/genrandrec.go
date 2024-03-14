package util

type GenRandField func() interface{}

// RandRec generates random record
type RandRec struct {
	Fields map[string]GenRandField
}

func (r *RandRec) Next() map[string]interface{} {
	ret := make(map[string]interface{})
	for k, v := range r.Fields {
		ret[k] = v()
	}
	return ret
}
