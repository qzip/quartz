module qz

go 1.16

require (
	github.com/arangodb/go-driver v1.6.0
	github.com/docker/go v1.5.1-1
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/pascaldekloe/jwt v1.10.0
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/crypto v0.5.0
	gopkg.in/square/go-jose.v2 v2.5.1
)

replace merkle => ../merkle
