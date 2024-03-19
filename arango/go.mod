module arango

go 1.19

require github.com/arangodb/go-driver v1.4.0

require (
	github.com/docker/go v1.5.1-1 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

require (
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/pkg/errors v0.9.1 // indirect
	qz v0.0.0
)

replace (
	merkle => ../ondc-server/merkle
	qz => ../ondc-server/qz
)
