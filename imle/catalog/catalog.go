package catalog

//Node  in the Catalog DAG (Distributed Acyclic Graph)
type Node struct {
	ID         string `json:"id"`
	UUID       string `json:"uuid,omitempty"`
	ParentUUID string `json:"parent,omitempty"`
	Namespace  string `json:"ns,omitempty"`
	Title      string `json:"title"`
	URL        string `json:"url,omitempty"`
	Tags       []Tag  `json:"tag,omitempty"`
}

//Tag help in searching & filtering
type Tag struct {
	Namespace string `json:"ns,omitempty"`
	Key       string `json:"key"`
	Value     string `json:"value,omitempty"`
}
