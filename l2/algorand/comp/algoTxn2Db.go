package comp

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"qz/seq"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type AlgoTxn2Db struct {
	DbConnection  string `json:"dbConnection"`
	DataInCtxName string `json:"confirmedTxnKey,omitempty"`
	//private fields
	helper  seq.CtxHelper
	errChan chan error
	db      *sql.DB
}

func (an *AlgoTxn2Db) SetCtxHelper(hlp seq.CtxHelper) {
	an.helper = hlp
}

func (an *AlgoTxn2Db) SetChanErr(ce chan error) {
	an.errChan = ce
}

const stmt = "REPLACE INTO cas (w3cdid, namespace, cas_data, tmstamp) VALUES(?,?,?,?)"

// implements commands.Pipeline interface
func (an *AlgoTxn2Db) Process(ctx context.Context) {
	an.helper.SetExecStatus(seq.ExSrunning)
	if an.DataInCtxName == "" {
		an.DataInCtxName = DataOutCtxName // from algoNotarize upstream
	}
	txn := an.helper.Value(an.DataInCtxName).(*AlgoTransaction)
	if txn == nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- fmt.Errorf("AlgoTxn2Db.Process: %s  is nil in helper context", an.DataInCtxName)
		return
	}
	if db, err := sql.Open("mysql", an.DbConnection); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return

	} else {
		an.db = db
	}
	defer an.db.Close()

	pstmt, err := an.db.PrepareContext(ctx, stmt)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return

	}
	defer pstmt.Close()

	p, err := json.MarshalIndent(txn, "\n", " ")
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	tmstamp := time.Now()
	if _, err = pstmt.ExecContext(ctx, txn.W3Cdid, txn.Namespace, p, tmstamp); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	an.helper.SetExecStatus(seq.ExSok)
}
