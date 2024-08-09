package catalog

// Node  in the Catalog DAG (Distributed Acyclic Graph)
type Node struct {
	W3cDid       string    `json:"w3cdid"`
	ParentW3cDid string    `json:"parent,omitempty"`
	Title        string    `json:"title"`
	URL          string    `json:"url,omitempty"`
	Tags         []Tag     `json:"tag,omitempty"`
	Specs        ProdSpecs `json:"specs,omitempty"`
}

// Tag help in searching & filtering
type Tag struct {
	Namespace string `json:"ns,omitempty"`
	Key       string `json:"key"`
	Value     string `json:"value,omitempty"`
}

type ProdSpecs struct {
	Namespace        string                 `json:"ns"` // <domain>:<schema> , the last : defines schema
	W3cDid           string                 `json:"w3cdid"`
	ProdAttributes   map[string]interface{} `json:"attributes,omitempty"`
	ComponentsW3cDid []string               `json:"components,omitempty"`
}
