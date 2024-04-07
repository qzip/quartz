package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

func main() {
	src := fmt.Sprintf("in.qzip.l2algo.v%v", time.Now())
	fmt.Println(src)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, []byte(src))

	fmt.Printf("%s\n", dst)

	i := new(big.Int)
	i.SetString(string(dst), 16) // hex

	fmt.Println(i.Text(24))

}
