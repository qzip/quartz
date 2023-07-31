package ops

import "context"

//Opration provides the Https JSON request/response functions
type Opration interface {
	Name() string
	Handler(ctx context.Context, req interface{}) (res interface{}, err error)
}

var opsReg = make(map[string]Opration)

func Register(op Opration) {
	opsReg[op.Name()] = op
}

func GetOps(name string) Opration {
	return opsReg[name]
}
