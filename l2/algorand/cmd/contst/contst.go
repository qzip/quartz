package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
)

const algodAddress = "http://127.0.0.1:8080"
const algodToken = "aa8cb03e14c1bfcd215512f520ee0ef6727d8876711a635f8b4f077f12ec6555"

func main() {
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("Issue with creating algod client: %s\n", err)
		return
	}
	status, err := algodClient.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("Error getting status: %s\n", err)
		return
	}
	statusJSON, err := json.MarshalIndent(status, "", "\t")
	if err != nil {
		fmt.Printf("Can not marshall status data: %s\n", err)
	}
	fmt.Printf("%s\n", statusJSON)
}
