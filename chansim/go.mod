module chansim

go 1.22.0

require qz v0.0.0

require (
	bc v0.0.0 // indirect
	github.com/cometbft/cometbft v0.38.5 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	merkle v0.0.0 // indirect
)

replace (
	bc => ../bc
	merkle => ../merkle
	qz => ../qz
)
