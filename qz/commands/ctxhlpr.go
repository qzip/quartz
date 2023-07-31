/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

*/

package commands

import (
	"context"
	"fmt"
	"time"
)

//ContextHelper wraps the parent context to extend the k/v map
// though WithValue func might suite one off kv setting,
// but adding multiple values will create a deep nesting of context
type ContextHelper struct {
	context.Context

	super context.Context
	kv    map[interface{}]interface{}
}

//NewContextHelper creates a wrapper
func NewContextHelper(parent context.Context) *ContextHelper {
	ch := &ContextHelper{
		super: parent,
		kv:    make(map[interface{}]interface{}),
	}
	return ch
}

//Deadline parent wrapper
func (ch *ContextHelper) Deadline() (deadline time.Time, ok bool) {
	return ch.super.Deadline()
}

//Done wrapper
func (ch *ContextHelper) Done() <-chan struct{} {
	return ch.super.Done()
}

//Err wrapper
func (ch *ContextHelper) Err() error {
	return ch.super.Err()
}

//Value checks if key/value pair present locally, else passes it up to parent context
func (ch *ContextHelper) Value(key interface{}) interface{} {
	if key == nil {
		return nil
	}
	if v, ok := ch.kv[key]; ok {
		return v
	}
	return ch.super.Value(key)

}

//KeyValMap returns the local key value map
func (ch *ContextHelper) KeyValMap() map[interface{}]interface{} {
	return ch.kv
}

func (ch *ContextHelper) String() string {
	return fmt.Sprintf("comman.ContextHelper {%v}", ch.kv)
}
