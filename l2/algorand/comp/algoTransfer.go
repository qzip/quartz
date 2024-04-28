package comp

import (
	"context"
	"crypto/ed25519"
	"crypto/sha512"
	"fmt"
	"qz/seq"
	"qz/util"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/mnemonic"
	"github.com/algorand/go-algorand-sdk/v2/transaction"
	"github.com/algorand/go-algorand-sdk/v2/types"
	canjson "github.com/docker/go/canonical/json"
)

const (
	w3CdidPrefix = "did:algorand:transfer:"
)

type AlgoTransfer struct {
	AlgoKeyMnemonic string      `json:"algoKeyMnemonic,omitempty"`
	AlgodAddress    string      `json:"algodAddress,omitempty"` // ~/Node/algod.net
	AlgodToken      string      `json:"algodToken,omitempty"`   // ~/Node/algod.token
	DataInCtxName   string      `json:"noteKey,omitempty"`
	AlgoTransferTo  string      `json:"algoTransferTo"`
	Amount          uint64      `json:"amountMicroAlgo"` //0.1 Algos = 100,000 microAlgos
	DataOutCtxName  string      `json:"confirmedTxnKey,omitempty"`
	W3CdidPrefix    string      `json:"w3cdidPrefix,omitempty"`
	Namespace       string      `json:"namespace,omitempty"`
	Metadata        interface{} `json:"metadata,omitempty"` // Note field will be canonial hash
	//private fields
	helper      seq.CtxHelper
	errChan     chan error
	algodClient *algod.Client
	pvtKey      ed25519.PrivateKey
	account     crypto.Account
}

func (an *AlgoTransfer) SetCtxHelper(hlp seq.CtxHelper) {
	an.helper = hlp
}

func (an *AlgoTransfer) SetChanErr(ce chan error) {
	an.errChan = ce
}

// https://github.com/algorand/go-algorand-sdk/blob/v2.4.0/examples/utils.go
// implements commands.Pipeline interface
func (an *AlgoTransfer) Process(ctx context.Context) {
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
		an.W3CdidPrefix = w3CdidPrefix
	}
	if an.Namespace == "" {
		an.Namespace = Namespace
	}

	if an.AlgoKeyMnemonic == "" {
		an.AlgoKeyMnemonic = AlgoKeyMnemonic
	}
	var err error

	if an.AlgoKeyMnemonic, err = util.ReplaceEnv(an.AlgoKeyMnemonic); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}

	if an.algodClient, err = algod.MakeClient(an.AlgodAddress, an.AlgodToken); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	if an.pvtKey, err = mnemonic.ToPrivateKey(an.AlgoKeyMnemonic); err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
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
		Metadata:  an.Metadata, // corelates the notarized txn with its context
	}

	/*p, err := json.MarshalIndent(txn, "\n", " ")
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}*/
	//fmt.Println(string(p))
	an.helper.SetKeyValue(an.DataOutCtxName, txn)
	an.helper.SetExecStatus(seq.ExSok)
}

func (an *AlgoTransfer) createTxn(ctx context.Context) (*types.Transaction, error) {

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
	var note []byte = nil
	if an.Metadata != nil {
		out, err := canjson.MarshalCanonical(an.Metadata)
		if err != nil {
			return nil, err
		}
		note = hash(out)
		//note = base64.StdEncoding.EncodeToString(h)
	}

	tx, err := transaction.MakePaymentTxn(
		an.account.Address.String(), an.AlgoTransferTo,
		0, note, "", //an.account.Address.String(),
		sp,
	)
	return &tx, err
}

// Hash returns SHA-2 SHA-512 hash
func hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}
