/*
Copyright (c) 2021, QzIP Blockchain Technology LLP

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

//Transactor low level transaction logic on account
// assumes higer level has JWX protection & auth validation
type Transactor struct {
	txnStore  TxnStore
	validator Validator
}

//TxnStore saves block to store
type TxnStore interface {
	SaveTxnBatch(txn *Batch) (*imle.Block, error)
}

//Validator validates the transaction
type Validator interface {
	Validate(batch *Batch) error
}

//NewTransactor instantiates a new transactor
func NewTransactor(txnStore TxnStore, validator Validator) *Transactor {
	tx := &Transactor{txnStore: txnStore, validator: validator}
	return tx
}

// Transact creates a txn block after validation
func (tx *Transactor) Transact(txn *Batch) (block *imle.Block, err error) {
	err = tx.validator.Validate(txn)
	if err == nil {
		block, err = tx.txnStore.SaveTxnBatch(txn)
	}
	return
}

//AmountValidator cascading basic validator
type AmountValidator struct {
	Root Validator
}

//Validate checks basic sanity
func (v *AmountValidator) Validate(batch *Batch) (err error) {
	if v.Root != nil {
		if err = v.Root.Validate(batch); err != nil {
			return
		}
	}
	err = v.validateAccounts(batch.Transactions)
	return
}

func (v *AmountValidator) validateAccounts(txns []Transaction) (err error) {
	var amt *Amount
	var bal int64
	for i, txn := range txns {
		if amt == nil {
			amt = &txn.TxnAmount
			bal = 0
		} else {
			if !amt.Compatible(&txn.TxnAmount) {
				err = fmt.Errorf("AmountValidator: account name:  %v in index %v not compatiable ", txn.TxnAccount.Key.String(), i)
				return
			}
			bal += txn.TxnAmount.Units
		}
	}
	if bal != 0 {
		err = fmt.Errorf("AmountValidator: does not balance by %v", bal)
	}
	return
}
