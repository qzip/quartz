package dlt

import (
	"imle"
	"time"
)

//TxnRef transaction reference
type TxnRef struct {
	TxnHash  imle.BHash `json:""`
	OutIndex int64      `json:""`
}

//VaultState UXTO (Consumed == false) state in the vault
type VaultState struct {
	StateURI string    `json:""`
	Ref      TxnRef    `json:""`
	Tmstamp  time.Time `json:""`
	Consumed bool      `json:""`
}

//ValidatedTransaction Taco runtime transaction
type ValidatedTransaction struct {
	TxnProposal       *Proposal
	TxnProposalMerkle imle.BHash       `json:""`
	TxnData           *TransactionData `json:""`
	TxnHash           imle.BHash
	TxnDataMerkle     imle.BHash
	ValidatorsSig     []OpaqueData
}

// Proposal transaction proposal
type Proposal struct {
	TxnData       *TransactionData `json:""`
	TxnDataMerkle imle.BHash
	ProposerSig   OpaqueData
	TimeToLiveSec int64
}

//TransactionData txn data whose hash represents TxnRef
// https://docs.corda.net/docs/corda-os/4.4/key-concepts-transactions.html#transaction-chains
//https://docs.corda.net/docs/corda-os/4.4/key-concepts-transactions.html#reference-states
type TransactionData struct {
	ContractHash    imle.BHash   `json:"contractHash"` //contract instance hash
	TxnName         string       `json:"txnName"`
	Tmstamp         time.Time    `json:"tmstamp"`
	InStates        []TxnRef     `json:"inStates"`
	OutStates       []OpaqueData `json:"outStates"`
	ReferenceStates []TxnRef     `json:"referenceStates"`
	Commands        []Command    `json:"commands"`
	Attachments     []OpaqueData `json:"attachments"`
}

//Command the command name and params if any
type Command struct {
	CmdName TxnCommand   `json:"cmdName"`
	CmdData []OpaqueData `json:"cmdData"`
}

//ContractState runtime contract state
type ContractState struct {
	ContractTemplateHash imle.BHash `json:"contractTemplateHash"`
	GenesisHash          imle.BHash `json:"genesisHash"`
	TxnHash              imle.BHash `json:"txnHash"`
	TxnMerkleRoot        imle.BHash `json:"txnMerkleRoot"`
	PrevStateHash        imle.BHash `json:"prevStateHash"`
	VaultMerkleRoot      imle.BHash `json:"vaultMerkleRoot"`
	Tmstamp              time.Time  `json:"tmstamp"`
}

//OpaqueData  opaque data
type OpaqueData struct {
	MetaData []Triple          `json:"metaData"`
	Data     imle.HashableData `json:"data"`
}

//Triple Type label Value
type Triple struct {
	DataType string            `json:"dataType"`
	Key      string            `json:"key"`
	Value    imle.HashableData `json:"value"`
}
