package blockchain

import "time"

//Namespace Namespace with metedata and data
type Namespace struct {
	Namespace   string                 `json:"namespace"`
	Name        string                 `json:"name"` // resource name
	Type        string                 `json:"type"` // resource type
	RefID       string                 `json:"refID"`
	ContentCas  string                 `json:"contentCas"`
	LastUpdated time.Time              `json:"lastUpdated"`
	MetaData    map[string]interface{} `json:"metadata"`
}
