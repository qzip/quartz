/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

*/

package commands

import (
	"context"
	"fmt"
	"qz/util"
	"reflect"
)

// HelperFactory component  extender
type HelperFactory interface {
	ComponentFactory
	CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (interface{}, error)
}

// CfgHelpersKey Global Helpers key name in config
var CfgHelpersKey = "helpers"

type PostInstall interface {
	InstallChildren(ctx context.Context, helper interface{}, param interface{}) error
}

// PostInstaller supports recursive helper installers as order of install is not guarnteed
type PostInstaller struct {
	postInstaller PostInstall
	param         interface{}
	helper        interface{} // needed to ensure helper factory can be reused multiple times
}

// BuildCtxHandlers Install Listner Components
// does not use err channel as it may not have been created at this point.
func BuildCtxHandlers(ctx context.Context, cfg map[string]interface{}) (context.Context, error) {
	var piArr []PostInstaller

	lst, ok := cfg[CfgHelpersKey]
	if !ok {
		return ctx, nil // no global listners
	}
	cfgFactArr, err := makeCfgCompFactArr(lst)
	if err != nil {
		return ctx, err
	}
	nuCtx := NewContextHelper(ctx)
	kv := nuCtx.KeyValMap()
	for _, cfa := range cfgFactArr {
		if cfa.Component == "" || cfa.CtxName == "" {
			continue
		}
		cf, ok := LookUpComponentFactory(cfa.Component)
		if !ok {
			return ctx, fmt.Errorf("commands.InstallGlobalHandlers: %v component not registered", cfa.Component)
		}
		util.DebugInfo(ctx, fmt.Sprintf("commands.InstallGlobalHandlers: lookup %v, got: %v\n", cfa.Component, reflect.TypeOf(cf)))
		helperFact, ok := cf.(HelperFactory)

		if !ok || helperFact == nil {
			return ctx, fmt.Errorf("commands.InstallGlobalHandlers: %v component is not of type commands.HelperFactory", cfa.Component)
		}

		// avoid creating nested helpers as sequence is not gauranteed
		helper, err := helperFact.CreateHelper(nuCtx, cfa.Param, cfg)
		if err != nil {
			util.DebugInfo(ctx, fmt.Sprintf("commands.InstallGlobalHandlers: error create helper %v %v\n", cfa.CtxName, err.Error()))
			return ctx, err
		}
		if pi, ok := cf.(PostInstall); ok {
			p := PostInstaller{postInstaller: pi, param: cfa.Param, helper: helper}
			piArr = append(piArr, p)
		}
		kv[cfa.CtxName] = helper
		util.DebugInfo(ctx, fmt.Sprintf("Installed Helper: %v, %v\n", cfa.CtxName, helper))

	} // for config factory array
	if len(piArr) > 0 {
		for _, pa := range piArr {
			if err = pa.postInstaller.InstallChildren(nuCtx, pa.helper, pa.param); err != nil {
				return nil, err
			}
		}
	}
	return nuCtx, nil
}
