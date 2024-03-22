/*
Copyright (c) 2020, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>
*/
package main

import (
	"encoding/base32"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%v <file name to compute sha> \n", os.Args[0])
		return
	}
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//SHAKE-256
	// A hash needs to be 64 bytes long to have 256-bit collision resistance.
	h := make([]byte, 64)
	// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum256(h, buf)
	//	fmt.Printf("SHAKE-256:%x\n", h)
	fmt.Println(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(h))

}
