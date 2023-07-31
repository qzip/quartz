package comp

//CfgComponent standard way of defining component configuration
type CfgComponent struct {
	Component string      `json:"component"`
	Param     interface{} `json:"param"`
}
