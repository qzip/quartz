package comp

import (
	"context"
	"l2/algorand/util"
	"qz/seq"
)

type Md5HashFile struct {
	DataOutCtxName  string      `json:"noteKey,omitempty"`
	MetadataCtxName string      `json:"metadataKey,omitempty"`
	FileName        string      `json:"filename,omitempty"`
	Metadata        interface{} `json:"metadata,omitempty"`
	//private fields
	helper  seq.CtxHelper
	errChan chan error
}

func (an *Md5HashFile) SetCtxHelper(hlp seq.CtxHelper) {
	an.helper = hlp
}

func (an *Md5HashFile) SetChanErr(ce chan error) {
	an.errChan = ce
}

// implements commands.Pipeline interface
func (an *Md5HashFile) Process(ctx context.Context) {
	an.helper.SetExecStatus(seq.ExSrunning)
	if an.DataOutCtxName == "" {
		an.DataOutCtxName = DataInCtxName // for Notarize downstream pipeline
	}
	md5, err := util.AssetMetadataHashFromFile(an.FileName)
	if err != nil {
		an.helper.SetExecStatus(seq.ExSerror)
		an.errChan <- err
	}
	md5Hex := "md5-hex:" + md5
	an.helper.SetKeyValue(an.DataOutCtxName, md5Hex)
	if len(an.MetadataCtxName) > 0 && an.Metadata != nil {
		an.helper.SetKeyValue(an.MetadataCtxName, an.Metadata)
	}
	an.helper.SetExecStatus(seq.ExSok)
}
