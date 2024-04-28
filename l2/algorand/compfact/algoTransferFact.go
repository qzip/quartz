package compfact

import (
	"context"
	"encoding/json"
	"l2/algorand/comp"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type AlgoTransferFact struct {
}

// implements qz/command/ComponentFactory
func (anf *AlgoTransferFact) Name() string {
	return "comp.AlgoTransferFactory"
}

// implements qz/command/ComponentFactory
func (anf *AlgoTransferFact) Help() string {
	help := `

= Devnet

 innomon@raspberrypi:~/sandbox/aloagri/algo-devnode $ goal network create --devMode
  
  innomon@raspberrypi:~/sandbox/aloagri/algo-devnode $ goal network start -d ./Node
 
  goal network start -d ./Node
  goal network stop -d ./Node
  
   export ALGORAND_DATA="/home/innomon/sandbox/aloagri/algo-devnode/Node"
 
 goal account list -d ./Node
 goal account export -a {address_from_an_online_account_from_above_command_output}
 
 innomon@raspberrypi:~/sandbox/aloagri/algo-devnode $ goal account export -a 5PJW2RFEUVNEONMGNUKHZR6LBRQM6XMZQ4KV2FMWQYDOD4UCKCG7MQFHZU -d ./Node
 Exported key for account 5PJW2RFEUVNEONMGNUKHZR6LBRQM6XMZQ4KV2FMWQYDOD4UCKCG7MQFHZU: "word social expire glory unknown embark vital sing flash lonely ready sentence series now thrive unveil arrest nominee brass seminar federal minor north abandon lawsuit"
 
 export ALGO_MNEMONIC="word social expire glory unknown embark vital sing flash lonely ready sentence series now thrive unveil arrest nominee brass seminar federal minor north abandon lawsuit"
 
	
	`
	return help
}

// implements qz/command/ComponentFactory
func (anf *AlgoTransferFact) ComponentType() reflect.Type {
	return reflect.TypeOf(anf)
}

// Create implements seq.Runner interface method
func (anf *AlgoTransferFact) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	an := comp.AlgoTransfer{}
	by, err := json.Marshal(param)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	err = json.Unmarshal(by, &an)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	an.SetCtxHelper(helper)
	an.SetChanErr(errCh)

	return &an
}
