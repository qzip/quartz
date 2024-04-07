module l2

go 1.22.1

require github.com/algorand/go-algorand-sdk/v2 v2.4.0

require (
	github.com/cometbft/cometbft v0.38.5 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

require (
	bc v0.0.0
	github.com/algorand/go-codec/codec v1.1.10 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	merkle v0.0.0
)

replace (
	bc => ../bc
	merkle => ../merkle
	qz => ../qz
)
