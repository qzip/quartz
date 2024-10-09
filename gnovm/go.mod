module gnovm

go 1.22.1

require (
	bc v0.0.0
	github.com/go-sql-driver/mysql v1.8.1
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.5 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/gnolang/overflow v0.0.0-20170615021017-4d914c927216 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/cometbft/cometbft v0.38.5 // indirect
	github.com/gnolang/gno v0.2.0
	golang.org/x/crypto v0.25.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	merkle v0.0.0
)

replace (
	bc => ../bc
	merkle => ../merkle
)
