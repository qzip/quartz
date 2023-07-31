/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

*/

package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/event"
)

// CfgGlobalListnerKey Key name
// array of listners :
var CfgGlobalListnerKey = "listners"

// EventListnerFactory component  extender
type EventListnerFactory interface {
	ComponentFactory
	CreateEventListner(ctx context.Context, param interface{}, cfg map[string]interface{}) (event.EventListner, error)
}

// CfgCompFact config param for Event Listner Factory
type CfgCompFact struct {
	Component string      `json:"component"`
	CtxName   string      `json:"name,omitempty"`
	Param     interface{} `json:"param,omitempty"`
}

// InstallGlobalListners Install Listner Components
// does not use err channel as it may not have been created at this point.
func InstallGlobalListners(ctx context.Context, cfg map[string]interface{}) error {
	lst, ok := cfg[CfgGlobalListnerKey]
	if !ok {
		return nil // no global listners
	}
	cfgFactArr, err := makeCfgCompFactArr(lst)
	if err != nil {
		return err
	}
	evtBus := event.EvtBusFromContext(ctx)

	for _, cfa := range cfgFactArr {
		if cfa.Component == "" {
			continue
		}
		cf, ok := LookUpComponentFactory(cfa.Component)
		if !ok {
			return fmt.Errorf("commands.InstallGlobalListners: %v component not registered", cfa.Component)
		}
		lstFact, ok := cf.(EventListnerFactory)
		if !ok {
			return fmt.Errorf("commands.InstallGlobalListners: %v componet not a EventListnerFactory", cfa.Component)
		}
		evtListner, err := lstFact.CreateEventListner(ctx, cfa.Param, cfg)
		if err != nil {
			return err
		}
		evtBus.Subscribe(evtListner)

	} // for config factory array
	return nil
}

func makeCfgCompFactArr(params interface{}) ([]CfgCompFact, error) {
	by, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	cfgArr := make([]CfgCompFact, 0)
	err = json.Unmarshal(by, &cfgArr)
	if err != nil {
		return nil, err
	}
	return cfgArr, nil
}
