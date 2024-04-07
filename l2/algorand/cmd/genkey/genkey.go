package main

import (
	"encoding/base64"
	"log"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	account := crypto.GenerateAccount()
	mn, err := mnemonic.FromPrivateKey(account.PrivateKey)

	if err != nil {
		log.Fatalf("failed to generate account: %s", err)
	}

	log.Printf("Address: %s\n", account.Address)
	log.Printf("Mnemonic: %s\n", mn)
	seed := account.PrivateKey.Seed()
	b64 := base64.StdEncoding.EncodeToString([]byte(seed))
	log.Printf("base64 private key: %s\n", b64)

}
