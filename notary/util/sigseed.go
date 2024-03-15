package util

import (
	"crypto"
	"fmt"
	"os"

	"crypto/ed25519"
)

// change seed with every version
// NOTE: Always load seed from env :
// export SIG_SEED = <string genererated by exec/genseed>
// exec/genseed
// version: in.qzip.l2algo.v2021-04-02 13:20:16.085180202 +0530 IST m=+0.000108926
var Seed = []byte("f86fb9hhgh744jan4gag8a482gl7c46m469j57a7dn5dd8i6akeegfjie003660lfk5gg68n0g30i22nhdj71da8dkilmlemnk9fdam8624mb36jbgg6c97hkm")

func LoadSeed() (err error) {
	return LoadSeedEnv("SIG_SEED")
}

func LoadSeedEnv(envar string) (err error) {
	mn := os.Getenv(envar)
	if mn == "" {
		err = fmt.Errorf("env var not set, export %s=\"seed random\"", envar)
		return
	}
	Seed = []byte(mn)
	return nil
}
func GetSuperUserKey() ed25519.PrivateKey {
	// Generate an ed25519 keypair. This should never fail
	return ed25519.NewKeyFromSeed(Seed)
}

func GetUserKey(mobile, mmid string) ed25519.PrivateKey {
	hash := crypto.SHA512.New()
	hash.Write([]byte(mobile))
	hval := hash.Sum(nil)
	suKey := GetSuperUserKey()
	sig := ed25519.Sign(suKey, hval)
	return ed25519.NewKeyFromSeed(sig)
}
