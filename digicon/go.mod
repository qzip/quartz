module digicon

go 1.22.1

require (
	github.com/docker/go v1.5.1-1
	imle v0.0.0
)

require (
	github.com/cometbft/cometbft v0.37.2 // indirect
	merkle v0.0.0 // indirect
)

replace (
	imle => ../imle
	merkle => ../merkle
)
