package main

import (
	"log"

	"github.com/algorand/go-algorand-sdk/v2/crypto"
	"github.com/algorand/go-algorand-sdk/v2/mnemonic"
)

func main() {
	account := crypto.GenerateAccount()
	mn, err := mnemonic.FromPrivateKey(account.PrivateKey)

	if err != nil {
		log.Fatalf("failed to generate account: %s", err)
	}

	log.Printf("Address: %s\n", account.Address)
	log.Printf("Mnemonic: %s\n", mn)
}
