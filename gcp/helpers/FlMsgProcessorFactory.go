package helpers

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"alo/digicon"

	"gocloud.dev/pubsub"
)

//FlMsgProcessorFactory GCP Pub Sub Message to File
type FlMsgProcessorFactory struct {
}

//CreateHelper creates a cmd.MsgConsumerHelper Object
func (fs *FlMsgProcessorFactory) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (hlpr interface{}, err error) {
	fl, ok := param.(string)
	if !ok {
		return nil, fmt.Errorf("FlMsgProcessorFactory.CreateProcessor: %v not a string, expecting filename", param)

	}
	if ok := strings.HasSuffix(fl, "/"); !ok {
		fl = fl + "/"
	}
	hlpr = &Msg2FileHelper{fl: fl}
	return
}

//Msg2FileHelper implements cmd.MsgConsumerHelper saves msessage to file
type Msg2FileHelper struct {
	fl string
}

//Open implements cmd.MsgConsumerHelper method
func (m2f *Msg2FileHelper) Open(ctx context.Context) error {
	// nothing to open, dummy call for interface implementation
	return nil
}

//Consume implements cmd.MsgConsumerHelper method
func (m2f *Msg2FileHelper) Consume(ctx context.Context, m *pubsub.Message) (err error) {
	msg2fl := make(map[string]interface{}, 0)
	if m.Metadata != nil {
		msg2fl["meta"] = m.Metadata
	}
	data := make(map[string]interface{}, 0)
	if er := json.Unmarshal(m.Body, &data); er != nil {
		msg2fl["data"] = m.Body
	} else {
		msg2fl["data"] = data
	}

	p, err := json.MarshalIndent(msg2fl, "\n", " ")
	if err != nil {
		return
	}
	fln := hex.EncodeToString(digicon.ComputeHash(msg2fl))

	if err = ioutil.WriteFile(m2f.fl+fln+".json", p, 0644); err != nil {
		return
	}
	return nil
}

//Close implements cmd.MsgConsumerHelper method
func (m2f *Msg2FileHelper) Close(ctx context.Context) {
	//dummy,nothing to close.
}

// component methods

// Name to register in the component registry
func (fs *FlMsgProcessorFactory) Name() string {
	return "helpers.FlMsgProcessorFactory"
}

// Help return help string
func (fs *FlMsgProcessorFactory) Help() string {
	return `
	  
	  For printing to a dir, specify the file path.
	  "param": "path/to/dir/"
	`
}

// ComponentType returns the component type
func (fs *FlMsgProcessorFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(&FlMsgProcessorFactory{})
}
