module l2

go 1.22.1

require (
	github.com/algorand/go-algorand-sdk v1.24.0
	github.com/algorand/go-algorand-sdk/v2 v2.4.0
	qz v0.0.0-00010101000000-000000000000
)

require (
	github.com/algorand/avm-abi v0.1.1 // indirect
	github.com/cometbft/cometbft v0.38.5 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

require (
	bc v0.0.0
	github.com/algorand/go-codec/codec v1.1.10 // indirect
	golang.org/x/crypto v0.21.0
	merkle v0.0.0
)

replace (
	bc => ../bc
	merkle => ../merkle
	qz => ../qz
)
