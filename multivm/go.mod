module multivm

go 1.20

require (
	github.com/dop251/goja v0.0.0-20230707174833-636fdf960de1
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca
	qz v0.0.0-00010101000000-000000000000
)

require (
	github.com/chzyer/readline v1.5.0 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/google/pprof v0.0.0-20230207041349-798e818bf904 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace (
	merkle => ../merkle
	qz => ../qz
)
