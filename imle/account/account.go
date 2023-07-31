/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

/*
Package account is a Account batch transaction model
*/
package account

import (
	"fmt"
	"imle"
)

//Batch top level struct, contained inside a ledger.Transaction
// an accounting transaction in a simple accounting system
// the batch is embedded in a ledger.Transaction
// this account.Transaction is specific to the accounting sub system
// and is different from a ledger.Transaction.
type Batch struct {
	BatchID      string        `xml:"id,attr" json:"id"`
	Documents    []imle.Trie   `xml:"documents" json:"documents,omitempty"`
	Transactions []Transaction `xmL:"transactions" json:"transactions"`
}

//Transaction represents an accounting  transaction unit
// this is different from ledger.Transaction
type Transaction struct {
	// https://pkg.go.dev/golang.org/x/text/currency?tab=doc#Amount
	// https://github.com/golang/text/blob/v0.3.2/currency/currency.go
	TxnAmount  Amount      `xml:"amount" json:"amount"`
	TxnAccount imle.Trie   `xml:"account" json:"account"`
	Narration  imle.Trie   `xml:"narration" json:"narration"`
	Documents  []imle.Trie `xml:"documents" json:"documents,omitempty"`
}

//Amount currency & unit / cents
type Amount struct {
	Currency  string `xml:"currency" json:"currency"`
	Units     int64  `xml:"units" json:"units"`
	Precision uint8  `xml:"precision" json:"precision"` // represents decimals for INR its 2
}

// Account holds the account details & balance
type Account struct {
	Balance Amount      `xml:"balance" json:"balance"`
	Name    imle.Trie   `xml:"name" json:"name"`
	Meta    []imle.Trie `xml:"meta,omitempty" json:"meta,omitempty"`
	Kyc     []imle.Trie `xml:"kyc,omitempty" json:"kyc,omitempty"`
}

//Compatible checks if computable
func (a *Amount) Compatible(with *Amount) bool {
	return (a.Currency == with.Currency && a.Precision == with.Precision)
}

//Compare copmares if Currency/Precision matches
func (a *Amount) Compare(with *Amount) (eq int, err error) {
	if a.Compatible(with) {
		if a.Units == with.Units {
			eq = 0
		} else if a.Units > with.Units {
			eq = 1
		} else {
			eq = -1
		}
	} else {
		err = fmt.Errorf("Currency/Precision mismatch")
	}
	return
}
