package comp

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"qz/seq"
	"qz/util"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"
)

const (
	AlgodAddress   = "http://localhost:4001"
	DataInCtxName  = "data"
	DataOutCtxName = "confirmedTxn"
	Namespace      = "quartz.l2.algorand.comp.AlgoTransaction"
	W3CdidPrefix   = "did:algorand:txn:"
)

type AlgoTransaction struct {
	W3Cdid    string                                `json:"w3cdid"`
	Namespace string                                `json:"namespace"`
	TxnRes    models.PendingTransactionInfoResponse `json:"txnRes"`
}

type AlgoNotarize struct {
	Base64PrivateKey string `json:"base64PrivateKey"`
	AlgodAddress     string `json:"algodAddress,omitempty"`
	AlgodToken       string `json:"algodToken,omitempty"`
	DataInCtxName    string `json:"noteKey,omitempty"`
	DataOutCtxName   string `json:"confirmedTxnKey,omitempty"`
	W3CdidPrefix     string `json:"w3cdidPrefix,omitempty"`
	Namespace        string `json:"namespace,omitempty"`

	//private fields
	helper      seq.CtxHelper
	errChan     chan error
	algodClient *algod.Client
	pvtKey      ed25519.PrivateKey
	account     crypto.Account
}

func (an *AlgoNotarize) SetCtxHelper(hlp seq.CtxHelper) {
	an.helper = hlp
}

func (an *AlgoNotarize) SetChanErr(ce chan error) {
	an.errChan = ce
}

// https://github.com/algorand/go-algorand-sdk/blob/v2.4.0/examples/utils.go
// implements commands.Pipeline interface
func (an *AlgoNotarize) Process(ctx context.Context) {
	an.helper.SetExecStatus(seq.ExSrunning)

	if an.AlgodAddress == "" {
		an.AlgodAddress = AlgodAddress
	}
	if an.DataInCtxName == "" {
		an.DataInCtxName = DataInCtxName
	}
	if an.DataOutCtxName == "" {
		an.DataOutCtxName = DataOutCtxName
	}
	if an.W3CdidPrefix == "" {
		an.W3CdidPrefix = W3CdidPrefix
	}
	if an.Namespace == "" {
		an.Namespace = Namespace
	}
	var err error
	if an.Base64PrivateKey, err = util.ReplaceEnv(an.Base64PrivateKey); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}

	if an.algodClient, err = algod.MakeClient(an.AlgodAddress, an.AlgodToken); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	if sDec, err := base64.StdEncoding.DecodeString(an.Base64PrivateKey); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	} else {
		an.pvtKey = ed25519.NewKeyFromSeed(sDec)
	}
	if an.account, err = crypto.AccountFromPrivateKey(an.pvtKey); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}

	// create txn
	tx, err := an.createTxn(ctx)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	// sign txn
	_, sptxn, err := crypto.SignTransaction(an.pvtKey, *tx)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	// send txn
	pendingTxID, err := an.algodClient.SendRawTransaction(sptxn).Do(ctx)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	// await result
	confirmedTxn, err := transaction.WaitForConfirmation(an.algodClient, pendingTxID, 4, ctx)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	// https://pkg.go.dev/github.com/algorand/go-algorand-sdk/v2@v2.4.0/client/v2/common/models#PendingTransactionInfoResponse

	txn := &AlgoTransaction{
		W3Cdid:    an.W3CdidPrefix + pendingTxID,
		Namespace: an.Namespace,
		TxnRes:    confirmedTxn,
	}
	an.helper.SetKeyValue(an.DataOutCtxName, *txn)
	an.helper.SetExecStatus(seq.ExSok)
}

func (an *AlgoNotarize) createTxn(ctx context.Context) (*types.Transaction, error) {
	note := an.helper.Value(an.DataInCtxName)
	sp, err := an.algodClient.SuggestedParams().Do(ctx)
	if err != nil {
		return nil, err
	}
	// check balance
	acctInfo, err := an.algodClient.AccountInformation(an.account.Address.String()).Do(ctx)
	if err != nil {
		return nil, err
	}
	util.DebugInfo(ctx, fmt.Sprintf("Account %s balance: %d microAlgos", an.account.Address.String(), acctInfo.Amount))

	if acctInfo.Amount < sp.MinFee {
		return nil, fmt.Errorf("algoNotarize.proess: acount balance is zero, please fund the account %s", an.account.Address.String())
	}
	tx, err := transaction.MakePaymentTxn(
		an.account.Address.String(), an.account.Address.String(),
		0, []byte(note.(string)), an.account.Address.String(),
		sp,
	)
	return &tx, err
}
