module fi

go 1.22.1

require (
	github.com/alecthomas/jsonschema v0.0.0-20210301060011-54c507b6f074
	merkle v0.0.0

)

require (
	github.com/cometbft/cometbft v0.37.2 // indirect
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0 // indirect
)

replace (
	merkle => ../merkle
	qz => ../qz
)
