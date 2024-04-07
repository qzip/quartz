/*
Copyright (c) 2019-21, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/util"
	"reflect"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

const addrKey = "algodAddress"
const toKey = "algodToken"

// Connectestconnection test
type Connectest struct {
}

func (ct *Connectest) Name() string {
	return reflect.TypeOf(*ct).String()
}

func (ct *Connectest) Help() string {
	return `

	{
	 "comment": "
		/* checks connection to the algorand node
		https://developer.algorand.org/docs/build-apps/setup/#3-run-your-own-node

		export ALGORAND_DATA=/var/lib/algorand
		cat $ALGORAND_DATA/algod.token
		cat $ALGORAND_DATA/algod.net
		*/ 
	 "

	 "component": "cmd.Connectest",

	 "algodAddress": "http://127.0.0.1:8080"
	 "algodToken" = "aa8cb03e14c1bfcd215512f520ee0ef6727d8876711a635f8b4f077f12ec6555"

	}
	`
}

func (ct *Connectest) Exec(ctx context.Context, cfg map[string]interface{}, errCh chan error) {
	util.DebugInfo(ctx, ct.Name()+" : START ")
	algodAddress := ""
	if addr, ok := cfg[addrKey]; ok {
		if s, ok := addr.(string); ok {
			algodAddress = s
		} else {
			err := fmt.Errorf("%v: algodAddress param not a string", ct.Name())
			ct.error(ctx, errCh, err)
			return
		}
	} else {
		err := fmt.Errorf("%v: algodAddress param not found", ct.Name())
		ct.error(ctx, errCh, err)
		return
	}
	var algodToken string
	if tok, ok := cfg[toKey]; ok {
		if s, ok := tok.(string); ok {
			algodToken = s
		} else {
			err := fmt.Errorf("%v: algodToken param not a string", ct.Name())
			ct.error(ctx, errCh, err)
			return
		}
	} else {
		err := fmt.Errorf("%v: algodToken param not found", ct.Name())
		ct.error(ctx, errCh, err)
		return
	}

	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		err := fmt.Errorf("%v: Issue with creating algod client: %s", ct.Name(), err)
		ct.error(ctx, errCh, err)
		return
	}
	status, err := algodClient.Status().Do(context.Background())
	if err != nil {
		err := fmt.Errorf("%v: Error getting status: %s", ct.Name(), err)
		ct.error(ctx, errCh, err)
		return
	}
	statusJSON, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		err := fmt.Errorf("%v: Can not marshall status data: %s", ct.Name(), err)
		ct.error(ctx, errCh, err)
		return
	}
	fmt.Printf("%s\n", statusJSON)
}

func (ct *Connectest) error(ctx context.Context, errCh chan error, err error) {
	util.DebugInfo(ctx, err.Error()+"\n")
	errCh <- commands.NewFatalError(err.Error() + "\n")
}
