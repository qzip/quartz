package main

import (
	"encoding/json"
	"fi/ledger"
	"fmt"
	"log"
	"time"

	"bc/blockchain"

	"github.com/alecthomas/jsonschema"
)

func main() {

	fmt.Println("= QzIP.in Financial Inclusion Accounting Blockchain JSON Schema")
	fmt.Println("version 0.1: Generated", time.Now())

	fmt.Print("\n\n\n")
	fmt.Println(voucherInfo)
	fmt.Print("\n\n.Voucher\n\n----\n\n")
	fmt.Println(genSchema(&ledger.Voucher{}))
	fmt.Print("\n\n----\n\n")

	fmt.Print("\n\n\n")
	fmt.Println(coaInfo)
	fmt.Print("\n\n.ChartOfAccounts\n\n----\n\n")
	fmt.Println(genSchema(&ledger.ChartOfAccounts{}))
	fmt.Print("\n\n----\n\n")

	fmt.Print("\n\n\n")
	fmt.Println(ledgerInfo)
	fmt.Print("\n\n.Ledger\n\n----\n\n")
	fmt.Println(genSchema(&ledger.Ledger{}))
	fmt.Print("\n\n----\n\n")

	fmt.Print("\n\n\n")
	fmt.Println(bcInfo)
	fmt.Print("\n\n.SignedBlockChain\n\n----\n\n")
	fmt.Println(genSchema(&blockchain.SignedBlockChain{}))
	fmt.Print("\n\n----\n\n")

}

func genSchema(in interface{}) string {
	schema := jsonschema.Reflect(in)
	json, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

var voucherInfo = `

The FI (Financial Inclusion) Accounting system, generates the vouchers via 
 events, during the data capture UI. All accounting entries (transactions) MUST be recorded via vouchers only.

None of the data should be deleted, not modified from backend system NOR vouchers should be deleted.

Apart from blockchain technical requirement, this condition is also mandatory as per the https://cleartax.in/s/gaap-india[Indian GAAP (Generally Accepted Accounting Principles )]

`
var coaInfo = `

The Chart of Accounts provides the hierarchy of the account ledger entries.
This needs to be provided just once and can be modified as well. It has no affect on blockchain
integrity, because its only used for report generation.

`

var ledgerInfo = `

The Voucher records are sequentialized and batched into blocks, which are chained cryptographically.
Thereafter, for each block generated, a ledger state is generated. This contains the account balances of
all accounts, applied from the valid vouchers, from current block till the first block.

Some vouchers may be rejected during this process, if it fails the validation rules.
The rejected vouchers are also recorded.

`

var bcInfo = `

Blockchain is a special case of txn chain. At a regular interval the blockchain process wakes up,
selects the transactions that occured between now and the last block; then creates a record:

Blocks are created and stored as txnchain (blockchain)

Signing is detached and save as seperate records. Thus a block can
have multiple signers

`
