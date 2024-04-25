package compfact

import (
	"context"
	"encoding/json"
	"l2/algorand/comp"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type AlgoTxn2DbFact struct {
}

// implements qz/command/ComponentFactory
func (anf *AlgoTxn2DbFact) Name() string {
	return "comp.AlgoTxn2DbFact"
}

// implements qz/command/ComponentFactory
func (anf *AlgoTxn2DbFact) Help() string {
	help := `
	Assumes Table structure as :
    CREATE TABLE IF NOT EXISTS cas (
		w3cdid	    VARCHAR(512) NOT NULL PRIMARY KEY,
		namespace	VARCHAR(512) DEFAULT ('com.aloagri.harvest2invoice'),
		cas_data	JSON,
		tmstamp	    TIMESTAMP
	);

	`
	return help
}

// implements qz/command/ComponentFactory
func (anf *AlgoTxn2DbFact) ComponentType() reflect.Type {
	return reflect.TypeOf(anf)
}

// Create implements seq.Runner interface method
func (anf *AlgoTxn2DbFact) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	an := comp.AlgoTxn2Db{}
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
