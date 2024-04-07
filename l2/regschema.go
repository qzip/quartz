package l2

import (
	"merkle"
	"reflect"
)

type SchemaValidator interface {
	Validate(schema, json []byte) error
}

type GenMerkleProofs interface {
	Hash() []byte
	SimpleProofsFromMap(json []byte) (rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string, err error)
}

type SchemaFactory interface {
	Name() string // registry
	Help() string
	ComponentType() reflect.Type
	//	Create(ctx context.Context, cfg map[string]interface{}, errChan chan error) (component interface{})
	Validator() SchemaValidator
	MerkleHasher() GenMerkleProofs
}

var regSchemaFactory = make(map[string]SchemaFactory)

// RegisterSchemaFactory register a component factory
func RegisterSchemaFactory(cfact SchemaFactory) {
	regSchemaFactory[cfact.Name()] = cfact
}

// LookUpSchemaFactory gets command from registry
func LookUpSchemaFactory(comp string) (cf SchemaFactory, ok bool) {
	cf, ok = regSchemaFactory[comp]
	return
}

// IterateSchemaFactory for each component call back
func IterateSchemaFactory(callBack func(SchemaFactory)) {
	for _, v := range regSchemaFactory {
		callBack(v)
	}
}
