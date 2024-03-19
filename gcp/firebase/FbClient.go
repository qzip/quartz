/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/util"
	"reflect"
	"strings"
	"time"

	app "firebase.google.com/go/v4"
	firebase "firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

// Global Config params keys (when ParamDocSubColSrc.ProjectID & CredPath) are absent
const (
	CfgParamProjectID    = "projectID"
	CfgParamCredPath     = "credentialsPath"
	CfgParamFirebasePfix = "firebase"
)

//FbMultiClient for usecase where more than 1 Firebase Db is required
type FbMultiClient struct {
}

//Name implements commands.Components Methods
func (hlpr *FbMultiClient) Name() string {
	return "firebase.MultiClient"
}

//Help implements commands.Components Methods
func (hlpr *FbMultiClient) Help() string {
	return `
	 FbMultiClient for usecase where more than 1 Firebase Db is required
	 Surrogate Helper for Firebase client: all parameters, except projectID are optional
	 "params": {
		"projectID": "alo-kheti-badi",
		"dabaseURL": "https://alo-kheti-badi.firebaseio.com/",
		"credPath" : "optional/path/to/file.json",
		"timeDurationSec" : 1,
	 }
	`
}

//ComponentType implements commands.Components Methods
func (hlpr *FbMultiClient) ComponentType() reflect.Type {
	return reflect.TypeOf(hlpr)
}

//CreateHelper implements commands.HelperFactory Methods
func (hlpr *FbMultiClient) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (fbClient interface{}, err error) {
	hx := &FbClientHelper{}
	return hx.CreateHelper(ctx, param, cfg)
}

//FbClientHelper firebase client
type FbClientHelper struct {
	fsClient *firebase.Client
	params   *ParamFsClientHelper
}

// ParamFsClientHelper configuration parameters
type ParamFsClientHelper struct {
	ProjectID string `json:"projectID,omitempty"`
	//DabaseURL is not specified then assumed to be the default database
	DabaseURL string `json:"dabaseURL,omitempty"`
	// CredPath is optional, or can be "default" or <path to cred file>
	CredPath string `json:"credentialsPath,omitempty"`
	//TimeDurationSec timer sleep interval while checking for done context
	TimeDurationSec time.Duration `json:"timeDurationSec,omitempty"`
}

//Name implements commands.Components Methods
func (hlpr *FbClientHelper) Name() string {
	return "firebase.Client"
}

//Help implements commands.Components Methods
func (hlpr *FbClientHelper) Help() string {
	return `
	 Helper for Firebase client: all parameters, except projectID are optional
	 "params": {
		"projectID": "alo-kheti-badi",
		"dabaseURL": "https://alo-kheti-badi.firebaseio.com/",
		"credPath" : "optional/path/to/file.json",
		"timeDurationSec" : 1,
	 }
	`
}

//ComponentType implements commands.Components Methods
func (hlpr *FbClientHelper) ComponentType() reflect.Type {
	return reflect.TypeOf(hlpr)
}

//CreateHelper implements commands.HelperFactory Methods
func (hlpr *FbClientHelper) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (fbClient interface{}, err error) {
	if hlpr.fsClient != nil {
		fbClient = hlpr.fsClient
		return
	}
	err = hlpr.open(ctx, param, cfg)
	if err == nil {
		var sleep time.Duration = time.Duration(1)
		if hlpr.params.TimeDurationSec > 1 {
			sleep = hlpr.params.TimeDurationSec
		}
		go func(ctx context.Context, sec time.Duration) {
			for {
				select {
				case <-ctx.Done():
					hlpr.close()
					return
				default:
					time.Sleep(sec * time.Second)
				}
			}
		}(ctx, sleep)
	}
	fbClient = hlpr.fsClient
	return
}

// https://firebase.google.com/docs/admin/setup/#go
func (hlpr *FbClientHelper) open(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	err := hlpr.getParams(ctx, param, cfg)
	if err != nil {
		return err
	}
	var appfb *app.App
	if hlpr.params.CredPath == "" || strings.EqualFold(hlpr.params.CredPath, "default") {
		if appfb, err = app.NewApp(ctx, nil); err != nil {
			return err
		}

	} else {
		opt := option.WithCredentialsFile(hlpr.params.CredPath)
		config := &app.Config{
			ProjectID:   hlpr.params.ProjectID,
			DatabaseURL: hlpr.params.DabaseURL,
		}
		if appfb, err = app.NewApp(ctx, config, opt); err != nil {
			return err
		}
	}
	if hlpr.params.DabaseURL == "" {
		if hlpr.fsClient, err = appfb.Database(ctx); err != nil {
			return err
		}
	} else {
		if hlpr.fsClient, err = appfb.DatabaseWithURL(ctx, hlpr.params.DabaseURL); err != nil {
			return err
		}
	}
	return nil
}
func (hlpr *FbClientHelper) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	/*if param == nil {
		return fmt.Errorf("FbClientHelper.open: nil param")
	}*/
	hlpr.params = &ParamFsClientHelper{}
	if param != nil {
		// get project ID
		by, err := json.Marshal(param)
		if err != nil {
			return err
		}
		err = json.Unmarshal(by, hlpr.params)
		if err != nil {
			return err
		}
	}
	if hlpr.params.ProjectID == "" {
		pid := util.FlattenMap(cfg, CfgParamFirebasePfix+"."+CfgParamProjectID)
		if pid == nil {
			return fmt.Errorf("FbClientHelper.open: firebase projectID param not specified")
		}
		var ok bool
		hlpr.params.ProjectID, ok = pid.(string)
		if !ok {
			return fmt.Errorf("FbClientHelper.open: firebase projectID param not a string")
		}
	}

	// check if credential path present
	if hlpr.params.CredPath == "" {
		pid := util.FlattenMap(cfg, CfgParamFirebasePfix+"."+CfgParamCredPath)
		if pid == nil {
			// assumed default
			hlpr.params.CredPath = ""

		}
		if p, ok := pid.(string); ok {
			hlpr.params.CredPath = p
		} /*else {
			return fmt.Errorf("FbClientHelper.open: firebase credential param not specified")
		}*/
	}

	return nil
}

func (hlpr *FbClientHelper) close() {
	if hlpr.fsClient != nil {
		//hlpr.fsClient.Close()
		hlpr.fsClient = nil
	}

}
