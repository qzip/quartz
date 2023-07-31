/*
Copyright (c) 2019-20, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

/*
Package event provides an event handling framework
*/
package event

import (
	"context"
	"fmt"
	"qz/util"
	"reflect"
)

// EventListner handles event
type EventListner interface {
	// IsDefault handle only if no other event listner subscribed
	IsDefault() bool
	// Filter listen for type of event, if nil then all events are passed
	Filter() []reflect.Type
	//
	OnEvent(ctx context.Context, evt interface{})
}
type EventBus interface {
	Subscribe(listener EventListner)
	//UnSubscribe(listener EventListner)
	Publish(ctx context.Context, evt interface{})
}

// SysEventBusKeyName this is the name for key in map/context
const SysEventBusKeyName = "SYS$EventBus"

// Default Implementation.
type EvtBus struct {
	globals []EventListner
	filters map[reflect.Type][]EventListner
}

// GlobalEventBus default event bus
var GlobalEventBus *EvtBus

func init() {
	GlobalEventBus = NewEvtBus()
}

// EvtBusFromContext gets the event bus from context if present, else returns the global bus
func EvtBusFromContext(ctx context.Context) (eb *EvtBus) {

	eb, ok := ctx.Value(SysEventBusKeyName).(*EvtBus)
	if !ok {
		eb = GlobalEventBus
	}
	return
}

// PublishFromContext publishes the event, from the event bus from context if present, else publishes to the global bus
func PublishFromContext(ctx context.Context, evt interface{}) (eb *EvtBus) {
	eb, ok := ctx.Value(SysEventBusKeyName).(*EvtBus)
	if !ok {
		eb = GlobalEventBus
	}
	eb.Publish(ctx, evt)
	return
}

// NewEvtBus creates a new instance of the Event Bus
func NewEvtBus() *EvtBus {
	return &EvtBus{
		globals: make([]EventListner, 0),
		filters: make(map[reflect.Type][]EventListner),
	}
}

// Subscribe add listner
func (eb *EvtBus) Subscribe(listener EventListner) {
	fArr := listener.Filter()
	if fArr == nil {
		eb.globals = append(eb.globals, listener)
	} else {
		for _, typ := range fArr {
			listners, ok := eb.filters[typ]
			if !ok {
				listners = make([]EventListner, 0)
			}
			eb.filters[typ] = append(listners, listener)
		}
	}
}

// Publish handles event publishing calls the subscribers
func (eb *EvtBus) Publish(ctx context.Context, evt interface{}) {
	eTyp := reflect.TypeOf(evt)
	util.DebugInfo(ctx, fmt.Sprintln("EvtBus.Publish: ", eTyp))
	handled := false
	var defaultListner EventListner = nil
	for k, lArr := range eb.filters {
		if k == eTyp {
			handled = true
			for _, l := range lArr {
				if l.IsDefault() {
					defaultListner = l
				} else {
					go l.OnEvent(ctx, evt)
				}

			}
		}

	} // for filters
	if !handled && defaultListner != nil {
		handled = true
		go defaultListner.OnEvent(ctx, evt)
	}
	for _, g := range eb.globals {
		if g.IsDefault() && handled {
			continue
		}
		go g.OnEvent(ctx, evt)
	} // for globals

}
