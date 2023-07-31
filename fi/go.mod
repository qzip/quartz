module fi

go 1.16

require (
	github.com/alecthomas/jsonschema v0.0.0-20210301060011-54c507b6f074
	merkle v0.0.0
	qz v0.0.0

)

replace (
	merkle => ../merkle
	qz => ../qz
)
