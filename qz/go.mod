module qz

go 1.21

toolchain go1.22.0

require (
	bc v0.0.0-00010101000000-000000000000
	github.com/arangodb/go-driver v1.6.0
	github.com/docker/go v1.5.1-1
	github.com/pascaldekloe/jwt v1.10.0
	golang.org/x/crypto v0.21.0
	gopkg.in/square/go-jose.v2 v2.6.0
)

require (
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/cometbft/cometbft v0.38.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	merkle v0.0.0 // indirect
)

replace (
	bc => ../bc
	merkle => ../merkle
)
