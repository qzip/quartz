module qz

go 1.16

require (
	bc v0.0.0-00010101000000-000000000000
	github.com/arangodb/go-driver v1.6.0
	github.com/docker/go v1.5.1-1
	github.com/pascaldekloe/jwt v1.10.0
	golang.org/x/crypto v0.5.0
	gopkg.in/square/go-jose.v2 v2.6.0
)

replace (
	bc => ../bc
	merkle => ../merkle
)
