module doltdb

go 1.21.6

require {
	qz v0.0.0
)

replace (
	merkle => ../ondc-server/merkle
	qz => ../ondc-server/qz
)
