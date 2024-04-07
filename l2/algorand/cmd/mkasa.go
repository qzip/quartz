/*
Copyright (c) 2019-21, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package cmd

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"l2/algorand/util"
	"qz/commands"
	"reflect"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	transaction "github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/types"
)

// AsaParams sub parameters for asset
type AsaParams struct {
	AssetName     string `json:"assetName"`
	UnitName      string `json:"unitName"`
	AssetURL      string `json:"assetURL"`
	MetadataHash  string `json:"metadataHash,omitempty"`
	DefaultFrozen bool   `json:"defaultFrozen,omitempty"`
	Decimals      uint32 `json:"decimals,omitempty"`
	TotalIssuance uint64 `json:"totalIssuance"`
	Note          string `json:"Note,omitempty"`
}

// CreateAsset command
type CreateAsset struct {
	algodAddress   string
	algodToken     string
	client         *algod.Client
	privateKey     ed25519.PrivateKey
	accountAddress types.Address
	asaParams      *AsaParams
	cfg            map[string]interface{}
}

func (ca *CreateAsset) Name() string {
	return reflect.TypeOf(*ca).String()
}

func (ca *CreateAsset) Help() string {
	return `
	{
	 "comment": "
		/* checks connection to the algorand node
		https://developer.algorand.org/docs/build-apps/setup/#3-run-your-own-node

		export  ALGO_KEY_PASSPHRASE="25 phrase secret"

		export ALGORAND_DATA=/var/lib/algorand
		cat $ALGORAND_DATA/algod.token
		cat $ALGORAND_DATA/algod.net
		*/
	    "	
	 "command": "cmd.CreateAsset",
      
	 "algodAddress": "http://127.0.0.1:8080",
	 "algodToken": "aa8cb03e14c1bfcd215512f520ee0ef6727d8876711a635f8b4f077f12ec6555",
	 "metadataText": "Text is converted to metadataHash, if metadataHash is empty", 
	 "metadataFile": "<file is used to generate hash if both metadataHash &  metadataText are empty>", 
	 "asa" : {
		"assetName":  "ONE" ,   
		"unitName":  "ONE" ,    
		"assetURL" :  "https://one.com" ,     
		"metadataHash"  : "32b5a06d6962bcbae6ca98123994004a",
		"defaultFrozen" : false,
		"decimals" :     0,
		"totalIssuance" : 100000000000,
		"note": "associated with the description of the asset"           
	 } 	
	}
	`
}

func (ca *CreateAsset) Exec(ctx context.Context, cfg map[string]interface{}, errCh chan error) {
	ca.cfg = cfg
	util.DebugInfo(ctx, ca.Name()+" : START ")

	if sk, err := util.GetPrivateKey(); err == nil {
		ca.privateKey = sk
	} else {
		ca.error(ctx, errCh, err)
		return
	}
	if err := ca.getClient(); err != nil {
		ca.error(ctx, errCh, err)
		return
	}
	ca.getAccountAddress()

	if err := ca.getAsaParams(); err != nil {
		ca.error(ctx, errCh, err)
		return
	}

	if err := ca.printStatus(); err != nil {
		ca.error(ctx, errCh, err)
		return
	}
	// Get network-related transaction parameters and assign
	txParams, err := ca.client.SuggestedParams().Do(context.Background())
	if err != nil {
		err := fmt.Errorf("%v: Error getting suggested tx params: %s", ca.Name(), err)
		ca.error(ctx, errCh, err)
		return
	}
	// Print asset info for newly created asset.
	util.PrettyPrint(txParams)

	//https://developer.algorand.org/docs/features/asa/
	// create asset
	txn, err := transaction.MakeAssetCreateTxn(
		ca.accountAddress.String(),
		[]byte(ca.asaParams.Note),
		txParams,
		ca.asaParams.TotalIssuance, ca.asaParams.Decimals,
		ca.asaParams.DefaultFrozen,
		ca.accountAddress.String(), ca.accountAddress.String(),
		ca.accountAddress.String(), ca.accountAddress.String(),
		ca.asaParams.UnitName,
		ca.asaParams.AssetName,
		ca.asaParams.AssetURL,
		ca.asaParams.MetadataHash,
	)
	if err != nil {
		err := fmt.Errorf("%v: Error MakeAssetCreateTxn: %s", ca.Name(), err)
		ca.error(ctx, errCh, err)
		return
	}
	// sign the transaction
	if txid, stx, err := crypto.SignTransaction(ca.privateKey, txn); err == nil {
		fmt.Printf("Transaction ID: %s\n", txid)
		sendResponse, err := ca.client.SendRawTransaction(stx).Do(context.Background())
		if err != nil {
			err := fmt.Errorf("%v: failed to send transaction: %s", ca.Name(), err)
			ca.error(ctx, errCh, err)
			return
		}
		fmt.Printf("Submitted transaction %s\n", sendResponse)
		// Wait for transaction to be confirmed
		util.WaitForConfirmation(txid, ca.client)
		// Retrieve asset ID by grabbing the max asset ID
		// from the creator account's holdings.
		act, err := ca.client.AccountInformation(ca.accountAddress.String()).Do(context.Background())
		if err != nil {
			err := fmt.Errorf("%v: failed to account info: %s", ca.Name(), err)
			ca.error(ctx, errCh, err)
			return
		}
		assetID := uint64(0)
		//	find newest (highest) asset for this account
		for _, asset := range act.CreatedAssets {
			if asset.Index > assetID {
				assetID = asset.Index
			}
		}
		// print created asset and asset holding info for this asset
		fmt.Printf("\nAsset ID: %d\n\n", assetID)
		printCreatedAsset(assetID, ca.algodAddress, ca.client)
		printAssetHolding(assetID, ca.algodAddress, ca.client)

	} else {
		errCh <- fmt.Errorf("%v: Error Signing Txn: %s", ca.Name(), err)
		return
	}

}

func (ca *CreateAsset) getAccountAddress() {
	publicKey := ca.privateKey.Public()
	cpk := publicKey.(ed25519.PublicKey)
	copy(ca.accountAddress[:], cpk[:])
}

func (ca *CreateAsset) printStatus() error {
	status, err := ca.client.Status().Do(context.Background())
	if err != nil {
		return fmt.Errorf("%v: Error getting status: %s", ca.Name(), err)

	}
	statusJSON, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		return fmt.Errorf("%v: Can not marshall status data: %s", ca.Name(), err)
	}
	fmt.Printf("%s\n\n", statusJSON)
	return nil
}

const (
	asaKey          = "asa"
	metadataTextKey = "metadataText"
	metadataFileKey = "metadataFile"
)

func (ca *CreateAsset) getAsaParams() error {
	if asa, ok := ca.cfg[asaKey]; ok {
		ca.asaParams = &AsaParams{}
		by, err := json.Marshal(asa)
		if err != nil {
			return err
		}
		err = json.Unmarshal(by, ca.asaParams)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%v: asa param not set", ca.Name())

	}
	// all good so far
	if ca.asaParams.MetadataHash == "" {
		if txt, ok := ca.cfg[metadataTextKey]; ok {
			if s, ok := txt.(string); ok {
				ca.asaParams.MetadataHash = util.AssetMetadataHash(s)
			} else {
				return fmt.Errorf("%v: %s param not a string", ca.Name(), metadataTextKey)

			}
		} else if fl, ok := ca.cfg[metadataFileKey]; ok {
			if s, ok := fl.(string); ok {
				var err error
				if ca.asaParams.MetadataHash, err = util.AssetMetadataHashFromFile(s); err != nil {
					return fmt.Errorf("%v: file name [%s] spcified by %s param %s ", ca.Name(), s, metadataFileKey, err.Error())
				}
			} else {
				return fmt.Errorf("%v: %s param not a string", ca.Name(), metadataTextKey)

			}
		}

	}
	return nil
}

func (ca *CreateAsset) getClient() error {

	if addr, ok := ca.cfg[addrKey]; ok {
		if s, ok := addr.(string); ok {
			ca.algodAddress = s
		} else {
			return fmt.Errorf("%v: algodAddress param not a string", ca.Name())

		}
	} else {
		return fmt.Errorf("%v: algodAddress param not found", ca.Name())
	}

	if tok, ok := ca.cfg[toKey]; ok {
		if s, ok := tok.(string); ok {
			ca.algodToken = s
		} else {
			return fmt.Errorf("%v: algodToken param not a string", ca.Name())

		}
	} else {
		return fmt.Errorf("%v: algodToken param not found", ca.Name())

	}

	algodClient, err := algod.MakeClient(ca.algodAddress, ca.algodToken)
	if err != nil {
		return fmt.Errorf("%v: Issue with creating algod client: %s", ca.Name(), err)

	}
	ca.client = algodClient
	return nil
}

// printCreatedAsset utility to print created assert for account
func printCreatedAsset(assetID uint64, account string, client *algod.Client) {

	act, err := client.AccountInformation(account).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	for _, asset := range act.CreatedAssets {
		if assetID == asset.Index {
			util.PrettyPrint(asset)
			break
		}
	}
}

// printAssetHolding utility to print asset holding for account
func printAssetHolding(assetID uint64, account string, client *algod.Client) {

	act, err := client.AccountInformation(account).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to get account information: %s\n", err)
		return
	}
	for _, assetholding := range act.Assets {
		if assetID == assetholding.AssetId {
			util.PrettyPrint(assetholding)
			break
		}
	}
}
func (ca *CreateAsset) error(ctx context.Context, errCh chan error, err error) {
	util.DebugInfo(ctx, err.Error()+"\n")
	errCh <- commands.NewFatalError(err.Error() + "\n")
}
