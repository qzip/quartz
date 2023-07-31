module ondc

go 1.18

require (
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	qz v0.0.0-00010101000000-000000000000
)

require (
	github.com/dlclark/regexp2 v1.4.1-0.20201116162257-a2a8dda75c91 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	github.com/dop251/goja v0.0.0-20220705101429-189bfeb9f530 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace (
	merkle => ../merkle
	qz => ../qz
)
