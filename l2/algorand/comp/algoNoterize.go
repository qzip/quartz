package comp

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"log"
	"qz/seq"
	"qz/util"

	"github.com/algorand/go-algorand-sdk/v2/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/v2/crypto"
)

const (
	AlgodAddress = "http://localhost:4001"
)

type AlgoNotarize struct {
	Base64PrivateKey string `json:"base64PrivateKey"`
	AlgodAddress     string `json:"AlgodAddress,omitempty"`
	AlgodToken       string `json:"AlgodToken,omitempty"`
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
	// check balance
	acctInfo, err := an.algodClient.AccountInformation(an.account.Address.String()).Do(ctx)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
		return
	}
	log.Printf("Account balance: %d microAlgos", acctInfo.Amount)
	// create txn
	// sign txn
	// send txn
	// await result

	an.helper.SetExecStatus(seq.ExSok)
}
