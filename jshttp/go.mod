module jshttp

go 1.18

require (
	golang.org/x/crypto v0.5.0 // indirect
	qz v0.0.0-00010101000000-000000000000
)

require (
	github.com/docker/go v1.5.1-1 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace (
	merkle => ../merkle
	qz => ../qz
)
