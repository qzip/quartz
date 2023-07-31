/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package event

import (
	"context"
	"log"
	"qz/util"
	"reflect"
)

// ExitEvent exit the system
type ExitEvent struct {
	Err error
}

type ExitEventListner struct {
}

func (l *ExitEventListner) IsDefault() bool {
	return true
}

// Filter listen for type of event, if nil then all events are passed
func (l *ExitEventListner) Filter() []reflect.Type {
	return nil
	// ft := make([]reflect.Type, 1)
	// ft[0] = reflect.TypeOf(&ExitEvent{})
	// return ft
	//make([]reflect.Type,1)

}

//OnEvent exit on ExitEvent
func (l *ExitEventListner) OnEvent(ctx context.Context, evt interface{}) {
	xe, ok := evt.(*ExitEvent)
	if ok {
		util.DebugInfo(ctx, xe.Err.Error())
		log.Fatal(xe.Err)
	}
}
