module gnovm

go 1.22.1

require (
	bc v0.0.0
	github.com/go-sql-driver/mysql v1.8.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/cometbft/cometbft v0.38.5 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	merkle v0.0.0 // indirect
)

replace (
	bc => ../bc
	merkle => ../merkle
)
