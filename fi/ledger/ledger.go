package ledger

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
	"merkle"
	"strconv"
	"time"
)

type Voucher struct {
	Header   VoucherHeader     `json:"header"`
	Entries  []VoucherEntry    `json:"entries"`
	Metadata map[string]string `json:"meta,omitempty"`
}

func (v Voucher) Hash() (hash []byte) {
	hash, _, _ = v.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (v Voucher) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	if len(v.Header.MerkleRoot) == 0 {
		me := make(map[string]merkle.Hasher)
		for i, t := range v.Entries {
			me["entries."+strconv.Itoa(i)] = t
		}
		v.Header.MerkleRoot, _, _ = merkle.SimpleProofsFromMap(me)
	}
	m := make(map[string]merkle.Hasher)
	m["header"] = v.Header

	for i, ele := range v.Entries {
		m["entries."+strconv.Itoa(i)] = ele
	}
	m["Metadata"] = HashedData(MapHasher(v.Metadata))
	return merkle.SimpleProofsFromMap(m)
}

type VoucherHeader struct {
	JwtAuth    []byte    `json:"jwt,omitempty"`
	Tmstamp    time.Time `json:"tmstamp,omitempty"`
	TxnRef     string    `json:"txnRef"`
	LedgerID   string    `json:"ledgerID"`
	TypeID     string    `json:"typeID"`
	MerkleRoot []byte    `json:"hash,omitempty"`
	Narration  string    `json:"narration,omitempty"`
}

//MerkleProof generates the merkle proof of the Data Log
func (v VoucherHeader) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["jwt"] = DataHasher(v.JwtAuth)
	m["tmstamp"] = HashedData(TimeHasher(v.Tmstamp))
	m["txnRef"] = DataHasher([]byte(v.TxnRef))
	m["ledgerID"] = DataHasher([]byte(v.LedgerID))
	m["typeID"] = DataHasher([]byte(v.TypeID))
	m["hash"] = HashedData(v.MerkleRoot)
	m["narration"] = DataHasher([]byte(v.Narration))

	return merkle.SimpleProofsFromMap(m)
}

func (v VoucherHeader) Hash() (hash []byte) {
	hash, _, _ = v.MerkleProof()
	return
}

type VoucherEntry struct {
	AccountID string `json:"account"`
	Amount    Amount `json:"amount"`
	Narration string `json:"narration,omitempty"`
}

func (v VoucherEntry) Hash() (hash []byte) {
	hash, _, _ = v.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (v VoucherEntry) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["account"] = DataHasher(v.AccountID)
	m["amount"] = v.Amount
	m["narration"] = DataHasher([]byte(v.Narration))

	return merkle.SimpleProofsFromMap(m)
}

//Amount currency & unit / cents
type Amount struct {
	Currency  string `xml:"currency" json:"currency"`
	Units     int64  `xml:"units" json:"units"`
	Precision uint8  `xml:"precision" json:"precision"` // represents decimals for INR its 2
}

func (v Amount) Hash() (hash []byte) {
	hash, _, _ = v.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (v Amount) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["currency"] = DataHasher(v.Currency)
	m["units"] = HashedData(IntHasher(v.Units))
	m["precision"] = HashedData(Uint8Hasher(v.Precision))

	return merkle.SimpleProofsFromMap(m)
}

type ChartOfAccounts struct {
	Nodes []CoaNode `json:"nodes"`
}

func (coa ChartOfAccounts) Hash() (hash []byte) {
	hash, _, _ = coa.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (coa ChartOfAccounts) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	for i, v := range coa.Nodes {
		m[strconv.Itoa(i)] = v
	}
	return merkle.SimpleProofsFromMap(m)
}

//AlgoAddress Algorand Address String
type AlgoAddress string

//Roles decide the allowed operations on the node
type Roles []string

//Hash implements merkle.Hasher
func (r Roles) Hash() []byte {
	hash := sha512.New()
	for _, v := range r {
		hash.Write([]byte(v))
	}
	return hash.Sum(nil)
}

type CoaNode struct {
	AddressRoles map[AlgoAddress]Roles `json:"addressRoles,omitempty"`
	Parent       string                `json:"parent,omitempty"`
	Name         string                `json:"name"`
	Currency     string                `json:"currency,omitempty"`
	IsLeaf       bool                  `json:"isLeaf,omitempty"`
}

func (coa CoaNode) Hash() (hash []byte) {
	hash, _, _ = coa.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (coa CoaNode) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["parent"] = DataHasher([]byte(coa.Parent))
	m["name"] = DataHasher([]byte(coa.Name))
	m["currency"] = DataHasher([]byte(coa.Currency))
	m["isLeaf"] = HashedData(BoolHasher(coa.IsLeaf))
	return merkle.SimpleProofsFromMap(m)
}

type LedgerTxnBlock struct {
	Tmstamp         time.Time          `json:"tmstamp"`
	LedgerID        string             `json:"ledgerID"`
	BalancesHash    []byte             `json:"hash"`
	AcceptedTxnsNdx []int64            `json:"acceptedTxnsNdx"`
	RejectedTxns    []RejectedBlockTxn `json:"rejectedTxns"`
}

func (lh LedgerTxnBlock) Hash() (hash []byte) {
	hash, _, _ = lh.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (lh LedgerTxnBlock) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["tmstamp"] = DataHasher(TimeHasher(lh.Tmstamp))
	m["ledgerID"] = DataHasher([]byte(lh.LedgerID))
	m["hash"] = HashedData(lh.BalancesHash)

	m["acceptedTxnsNdx"] = HashedData(IntArrHasher(lh.AcceptedTxnsNdx))

	for i, t := range lh.RejectedTxns {
		m[strconv.Itoa(i)] = t
	}

	return merkle.SimpleProofsFromMap(m)
}

type RejectedBlockTxn struct {
	BlockIndex   int64  `json:"blockIndex"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func (rt RejectedBlockTxn) Hash() (hash []byte) {
	hash, _, _ = rt.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (rt RejectedBlockTxn) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["blockIndex"] = HashedData(IntHasher(rt.BlockIndex))
	m["errorCode"] = DataHasher([]byte(rt.ErrorCode))
	m["errorMessage"] = DataHasher([]byte(rt.ErrorMessage))

	return merkle.SimpleProofsFromMap(m)
}

type AccountBalances map[string]Amount

//Hash implements merkle.Hasher
func (rt AccountBalances) Hash() (hash []byte) {
	hash, _, _ = rt.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (rt AccountBalances) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	for k, v := range rt {
		m[k] = v
	}
	return merkle.SimpleProofsFromMap(m)
}

type Ledger struct {
	Header      LedgerTxnBlock    `json:"header"`
	Balances    AccountBalances   `json:"balances"`
	BlockHash   []byte            `json:"blockHash,omitempty"`
	BlockHeight int64             `json:"blockHeight,omitempty"`
	Metadata    map[string]string `json:"meta,omitempty"`
}

//Hash implements merkle.Hasher
func (coa Ledger) Hash() (hash []byte) {
	hash, _, _ = coa.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (coa Ledger) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["header"] = coa.Header
	m["balances"] = coa.Balances
	m["blockHash"] = HashedData(coa.BlockHash)
	m["blockHeight"] = HashedData(IntHasher(coa.BlockHeight))
	m["Metadata"] = HashedData(MapHasher(coa.Metadata))
	return merkle.SimpleProofsFromMap(m)
}

type AccountBalance struct {
	AccountID string `json:"account"`
	Amount    Amount `json:"amount"`
}

//Hash implements merkle.Hasher
func (coa AccountBalance) Hash() (hash []byte) {
	hash, _, _ = coa.MerkleProof()
	return
}

//MerkleProof generates the merkle proof of the Data Log
func (coa AccountBalance) MerkleProof() (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) {
	m := make(map[string]merkle.Hasher)
	m["account"] = DataHasher([]byte(coa.AccountID))
	m["amount"] = coa.Amount

	return merkle.SimpleProofsFromMap(m)
}

//HashedData pass through wrapper
type HashedData []byte

//Hash implements merkle.Hasher
func (h HashedData) Hash() []byte {
	return h
}

//DataHasher data that partakes in Merkle Hashing
type DataHasher []byte

//Hash implements merkle.Hasher
func (h DataHasher) Hash() []byte {
	return Hash([]byte(h))
}

// Hash returns SHA-2 SHA-512 hash
func Hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}

func TimeHasher(t time.Time) []byte {
	hash := sha512.New()
	b, _ := t.MarshalBinary()
	hash.Write(b)
	return hash.Sum(nil)
}
func IntHasher(i int64) []byte {
	hash := sha512.New()
	hash.Write(Int2Bin(i))
	return hash.Sum(nil)
}
func IntArrHasher(iarr []int64) []byte {
	hash := sha512.New()
	for _, i := range iarr {
		hash.Write(Int2Bin(i))
	}
	return hash.Sum(nil)
}
func Int2Bin(i int64) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func Uint8Hasher(i uint8) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	hash := sha512.New()
	hash.Write(buf.Bytes())
	return hash.Sum(nil)
}
func BoolHasher(i bool) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, i); err != nil {
		panic(err)
	}
	hash := sha512.New()
	hash.Write(buf.Bytes())
	return hash.Sum(nil)
}

func MapHasher(m map[string]string) []byte {
	mx := make(map[string]merkle.Hasher)
	for k, v := range m {
		mx[k] = DataHasher(v)
	}
	hash, _, _ := merkle.SimpleProofsFromMap(mx)
	return hash
}
