module bc

go 1.21

toolchain go1.21.6

require (
	gopkg.in/square/go-jose.v2 v2.6.0
	merkle v0.0.0
)

require (
	github.com/cometbft/cometbft v0.38.5 // indirect
	golang.org/x/crypto v0.20.0 // indirect
)

replace merkle => ../merkle
