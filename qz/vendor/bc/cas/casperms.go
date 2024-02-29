package cas

// CasJSONPermission defines an access role for a namespace
// access roles <role> --> [field path]
type CasJSONPermission struct {
	Namespace   string              `json:"namespace"`
	Schema      string              `json:"schema"`
	AccessRoles map[string][]string `json:"roles"`
}
