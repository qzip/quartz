package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"imle"
	"merkle"
	"qz/commands"
	"reflect"

	"cloud.google.com/go/pubsub"
)

// QueueMsgProcessorFactory send message from queue to contract chain
type QueueMsgProcessorFactory struct {
}

// TransformMessage  transforms message to transaction logs
type TransformMessage interface {
	Transform(msg *pubsub.Message) (out *imle.MerkleDataLog, err error)
}

// MerkleDataLogSaver helper component should implement this interface
type MerkleDataLogSaver interface {
	SaveMerkleDataLog(*imle.MerkleDataLog) error
}

// MerkleProofPublisher helper component should implement this interface
type MerkleProofPublisher interface {
	PublishMerkleProof(rootHash []byte, proofs map[string]*merkle.SimpleProof, keys []string) error
}

// Closable if a helper implements this inteface the parent close method will Close children also
type Closable interface {
	Close(context.Context)
}

// MsgConsumer the helper that runs the message consumer pipeline
type MsgConsumer struct {
	contextKeys          CfgQueueMsgProcessorFactory
	TransformMessage     TransformMessage
	MerkleDataLogSaver   MerkleDataLogSaver
	MerkleProofPublisher MerkleProofPublisher
}

// CfgQueueMsgProcessorFactory pipleine context keys
type CfgQueueMsgProcessorFactory struct {
	TransformMessageKey     string `json:"transformMessageKey"`
	MerkleDataLogSaverKey   string `json:"merkleDataLogSaverKey"`
	MerkleProofPublisherKey string `json:"merkleProofPublisher,omitempty"`
}

// CreateHelper implements Helper Factory method
func (q2c *QueueMsgProcessorFactory) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (helper interface{}, err error) {
	by, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	cfgQue := &CfgQueueMsgProcessorFactory{}
	if err := json.Unmarshal(by, cfgQue); err != nil {
		return nil, err
	}
	helper = &MsgConsumer{contextKeys: *cfgQue}

	return
}

// Consume implements MsgConsumerHelper method
func (mc *MsgConsumer) Consume(ctx context.Context, m *pubsub.Message) error {
	return nil
}

// Open implements MsgConsumerHelper method
func (mc *MsgConsumer) Open(ctx context.Context) error {
	helper := ctx.Value(mc.contextKeys.TransformMessageKey)

	if helper == nil {
		return commands.NewFatalError(fmt.Sprintf("MsgConsumer.Open: TransformMessage %v context helper does not exist", mc.contextKeys.TransformMessageKey))
	}
	var ok bool
	if mc.TransformMessage, ok = helper.(TransformMessage); !ok {
		return commands.NewFatalError(fmt.Sprintf("MsgConsumer.Open:  %v context helper not of type TransformMessage", mc.contextKeys.TransformMessageKey))

	}
	helper = ctx.Value(mc.contextKeys.MerkleDataLogSaverKey)

	if helper == nil {
		return commands.NewFatalError(fmt.Sprintf("MsgConsumer.Open: MerkleDataLogSave %v context helper does not exist", mc.contextKeys.MerkleDataLogSaverKey))
	}

	if mc.MerkleDataLogSaver, ok = helper.(MerkleDataLogSaver); !ok {
		return commands.NewFatalError(fmt.Sprintf("MsgConsumer.Open:  %v context helper not of type MerkleDataLogSaver", mc.contextKeys.MerkleDataLogSaverKey))

	}
	//optional publisher
	if mc.contextKeys.MerkleProofPublisherKey == "" {
		return nil
	}
	helper = ctx.Value(mc.contextKeys.MerkleProofPublisherKey)

	if helper == nil {
		return commands.NewFatalError(fmt.Sprintf("MsgConsumer.Open: MerkleProofPublisherKey %v context helper does not exist", mc.contextKeys.MerkleProofPublisherKey))
	}

	if mc.MerkleDataLogSaver, ok = helper.(MerkleDataLogSaver); !ok {
		return commands.NewFatalError(fmt.Sprintf("MsgConsumer.Open:  %v context helper not of type MerkleProofPublisherKey", mc.contextKeys.MerkleProofPublisherKey))

	}

	return nil
}

// Close implements MsgConsumerHelper method
func (mc *MsgConsumer) Close(ctx context.Context) {

}

// component methods

// Name to register in the component registry
func (q2c *QueueMsgProcessorFactory) Name() string {
	return "helpers.FlMsgProcessorFactory"
}

// Help return help string
func (q2c *QueueMsgProcessorFactory) Help() string {
	return `
	  
	  For printing to a dir, specify the file path.
	  "param": "path/to/dir/"
	`
}

// ComponentType returns the component type
func (q2c *QueueMsgProcessorFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(&QueueMsgProcessorFactory{})
}
