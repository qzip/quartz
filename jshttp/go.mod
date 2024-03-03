module jshttp

go 1.21

toolchain go1.22.0

require (
	golang.org/x/crypto v0.20.0 // indirect
	qz v0.0.0-00010101000000-000000000000
)

require (
	github.com/Workiva/go-datastructures v1.1.1
	github.com/docker/go v1.5.1-1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

replace (
	merkle => ../merkle
	qz => ../qz
)
