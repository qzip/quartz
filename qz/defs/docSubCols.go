package defs

const (
	//DocID id
	DocID = "id"
	//DocPATH path
	DocPATH = "path"
)

// DocSubCols firebase document & sub columns
type DocSubCols struct {
	Doc     map[string]interface{}
	SubCols map[string]map[string]interface{}
}
