module multivm

go 1.20

require (
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca
	qz v0.0.0-00010101000000-000000000000
)

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/docker/go v1.5.1-1 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace (
	merkle => ../merkle
	qz => ../qz
)
