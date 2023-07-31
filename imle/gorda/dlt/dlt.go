package dlt

//TacoTemplate Taco is to Gorda, what CorDapp is to Corda.
type TacoTemplate struct {
	TemplateURI     string           `json:"templateURI"`
	Roles           []Role           `json:"roles"`
	States          []StateDef       `json:"states"`
	ValidatorVM     ValidatorVMType  `json:"validatorVM"`
	TransactionDefs []TransactionDef `json:"transactionDefs"`
	CodePath        CasURI           `json:"codePath,omitempty"`       // optional
	LegalTemplates  []CasURI         `json:"legalTemplates,omitempty"` // local cas alias/uri
}

//CasURI Content Addressable store URI, can refer to local zip file dir content
type CasURI string

//Role is assigned to a partner during instantiation
type Role string

//StateDef represents a template of a state on ledger facts
// state schema URI :
//   <TemplateURI>:state:<Name>
type StateDef struct {
	Name      string `json:"name"`
	SchemaURI string `json:"schemaURI"`
	Observers []Role `json:"observers"`
}

//TransactionDef Defines transaction
type TransactionDef struct {
	Name         string     `json:"name"`
	InStateName  []string   `json:"inStateName"`
	OutStateName []string   `json:"outStateName"`
	Command      TxnCommand `json:"command"`
	Proposers    []Role     `json:"proposers"`
	Validators   []Role     `json:"validators"`
	//ValidatorFuncName provides access to VM executed func.
	//format is [<dir path>:]<func> like, "/codes/validator.jar:Validate"
	//incase, the code path is provided at template config top level,
	// it is also appended. If just function name is provided, then the ":"
	// should be omitted.
	ValidatorFuncName string `json:"validatorFuncName"`
}

//TxnCommand Command State Commands
type TxnCommand string

/*

FlowDef commented as its not needed, since Transaction Def
has a proposer role introduced.
//FlowDef Flow defs.
// initiator/responder functions can be overridden during instantiation
type FlowDef struct {
	Name              string           `json:"name"`
	Transactions      []TransactionDef `json:"transactions"`
	Initiators        []Role           `json:"initiators"`
	Responders        []Role           `json:"responders"`
	InitatorFuncName  string           `json:"initatorFuncName"`
	ResponderFuncName string           `json:"responderFuncName"`
}
*/

//ValidatorVMType VM types implemented
type ValidatorVMType string

const (
	//CEL google common expression language
	CEL ValidatorVMType = "cel"
	//FWRAP pronounced like frappe
	FWRAP ValidatorVMType = "fwrap"
	//SCHEME functional language
	SCHEME ValidatorVMType = "scheme"
	//JAVA java vm
	JAVA ValidatorVMType = "java"
	//LUA vm
	LUA ValidatorVMType = "lua"
	//RUBY vm
	RUBY ValidatorVMType = "ruby"
	//JS (Javascript) vm
	JS ValidatorVMType = "js"
	//WASM web assembly
	WASM ValidatorVMType = "wasm" // https://github.com/go-interpreter/wagon
	//PLUGIN go plugin
	PLUGIN ValidatorVMType = "plugin"
)
