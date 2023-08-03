package multivm

import "time"

/**
VM Script serialized in binary format
Record:
Hash, time stamp, vm type, file name, content

*/

type VM string

const (
	JS      VM = "js"       // https://github.com/dop251/goja
	WASM    VM = "wasm"     // https://github.com/tetratelabs/wazero
	CEL     VM = "cel"      //https://github.com/google/cel-go
	JOKER   VM = "joker"    // https://github.com/candid82/joker
	GNO     VM = "gno"      // https://github.com/gnolang/gno/tree/master/gnovm
	SKYLARK VM = "starlark" // https://github.com/google/starlark-go/blob/master/doc/spec.md
)

type VMScipt struct {
	Filename string
	VMtype   VM
	Tmstamp  time.Time
	Content  []byte
}
