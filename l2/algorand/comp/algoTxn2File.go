package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"qz/seq"
	"strings"
)

type AlgoTxn2File struct {
	DataInCtxName string `json:"confirmedTxnKey,omitempty"`
	OutFileName   string `json:"outFileName"`
	//private fields
	helper  seq.CtxHelper
	errChan chan error
	w       io.Writer
}

func (an *AlgoTxn2File) SetCtxHelper(hlp seq.CtxHelper) {
	an.helper = hlp
}

func (an *AlgoTxn2File) SetChanErr(ce chan error) {
	an.errChan = ce
}

// implements commands.Pipeline interface
func (an *AlgoTxn2File) Process(ctx context.Context) {
	an.helper.SetExecStatus(seq.ExSrunning)
	if an.DataInCtxName == "" {
		an.DataInCtxName = DataOutCtxName // from algoNotarize upstream
	}
	txn := an.helper.Value(DataInCtxName) //.(*AlgoTransaction)
	if txn == nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- fmt.Errorf("AlgoTxn2Db.Process: %s  is nil in helper context", an.DataInCtxName)
		return
	}
	p, err := json.MarshalIndent(txn, "\n", " ")
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	an.open()
	f, ok := (an.w).(*os.File)
	if ok {
		defer f.Close()
	}
	if _, er := an.w.Write(p); er != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- er
		return
	}
	an.helper.SetExecStatus(seq.ExSok)
}

func (an *AlgoTxn2File) open() {
	if an.OutFileName == "" {
		an.w = os.Stdout
	} else if strings.EqualFold(an.OutFileName, "stdio") {
		an.w = os.Stdout
	} else {
		if w, err := os.OpenFile(an.OutFileName, os.O_RDWR|os.O_CREATE, 0755); err != nil {
			an.w = os.Stdout
		} else {
			an.w = w
		}
	}

}
