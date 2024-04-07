package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

func WaitForConfirmation(txID string, client *algod.Client) (err error) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		//	fmt.Printf("error getting algod status: %s\n", err)
		return
	}
	lastRound := status.LastRound
	for {
		pt, _, erx := client.PendingTransactionInformation(txID).Do(context.Background())
		if erx != nil {
			//	fmt.Printf("error getting pending transaction: %s\n", err)
			err = erx
			return
		}
		if pt.ConfirmedRound > 0 {
			//	fmt.Printf("Transaction "+txID+" confirmed in round %d\n", pt.ConfirmedRound)
			break
		}
		fmt.Printf("waiting for confirmation\n")
		lastRound++
		status, err = client.StatusAfterBlock(lastRound).Do(context.Background())
	}
	return
}

//https://golang.org/pkg/crypto/md5/

func AssetMetadataHash(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	md5hash := h.Sum(nil)
	return hex.EncodeToString(md5hash)
}

func AssetMetadataHashFromFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	md5hash := h.Sum(nil)
	return hex.EncodeToString(md5hash), nil
}

// PrettyPrint prints Go structs
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
