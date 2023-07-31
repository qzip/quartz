/*
Copyright (c) 2021, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package helpers

import "reflect"

// SrvCas implements qz/commands.HelperFactory
type SrvCas struct {
}

//Name implements qz/commands.HelperFactory
func (sc *SrvCas) Name() string {
	return reflect.TypeOf(sc).Elem().String()
}

//Help implements qz/commands.HelperFactory
func (sc *SrvCas) Help() string {
	return `
	"component": "qz/helpers.SrvCas",
	"param": {
		"path": "HandlerHelperFactCtxName"
	}
	`
}

//ComponentType implements qz/commands.HelperFactory
func (sc *SrvCas) ComponentType() reflect.Type {
	return reflect.TypeOf(sc)
}
