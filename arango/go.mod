module arango

go 1.21

toolchain go1.22.1

require (
	github.com/arangodb/go-driver v1.6.0
	github.com/google/uuid v1.4.0
)

require (
	github.com/cometbft/cometbft v0.38.5 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	merkle v0.0.0 // indirect
)

require (
	bc v0.0.0 // indirect
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/pkg/errors v0.9.1 // indirect
	qz v0.0.0
)

replace (
	bc => ../bc
	merkle => ../merkle
	qz => ../qz
)
