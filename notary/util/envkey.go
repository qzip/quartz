/*
Copyright (c) 2019-21, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package util

import (
	"crypto/ed25519"
	"fmt"
	"os"

	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func GetPrivateKey() (sk ed25519.PrivateKey, err error) {
	return EnvGetPrivateKey("ALGO_KEY_PASSPHRASE")

}
func EnvGetPrivateKey(envar string) (sk ed25519.PrivateKey, err error) {
	mn := os.Getenv(envar)
	if mn == "" {
		err = fmt.Errorf("env var not set, export %s=\"25 phrase secret\"", envar)
		return
	}
	sk, err = mnemonic.ToPrivateKey(mn)
	return
}
