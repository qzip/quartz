module doltdb

go 1.21.6

require (
	qz v0.0.0
	bc v0.0.0
)

replace (
	merkle => ../merkle
	qz => ../qz
	bc => ../bc
)
