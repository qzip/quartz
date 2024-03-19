/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/util"
	"reflect"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// Global Config params keys (when ParamDocSubColSrc.ProjectID & CredPath) are absent
const (
	CfgParamProjectID    = "projectID"
	CfgParamCredPath     = "credentialsPath"
	CfgParamFirebasePfix = "firebase"
)

//FsClientHelper firebase client
type FsClientHelper struct {
	fsClient *firestore.Client
	params   *ParamFsClientHelper
}

// ParamFsClientHelper configuration parameters
type ParamFsClientHelper struct {
	ProjectID string `json:"projectID,omitempty"`
	// CredPath is optional, or can be "default" or <path to cred file>
	CredPath string `json:"credentialsPath,omitempty"`
	//TimeDurationSec timer sleep interval while checking for done context
	TimeDurationSec time.Duration `json:"timeDurationSec,omitempty"`
}

//Name implements commands.Components Methods
func (hlpr *FsClientHelper) Name() string {
	return "firestore.Client"
}

//Help implements commands.Components Methods
func (hlpr *FsClientHelper) Help() string {
	return `
	 Helper for Firebase client: all parameters, except projectID are optional
	 "params": {
		"projectID" : "my-project-id",
		"credPath" : "optional/path/to/file.json",
		"timeDurationSec" : 1,
	 }
	`
}

//ComponentType implements commands.Components Methods
func (hlpr *FsClientHelper) ComponentType() reflect.Type {
	return reflect.TypeOf(hlpr)
}

//CreateHelper implements commands.HelperFactory Methods
func (hlpr *FsClientHelper) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (fbClient interface{}, err error) {
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

func (hlpr *FsClientHelper) open(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	err := hlpr.getParams(ctx, param, cfg)
	if err != nil {
		return err
	}
	if hlpr.params.CredPath == "" || strings.EqualFold(hlpr.params.CredPath, "default") {
		if hlpr.fsClient, err = firestore.NewClient(ctx, hlpr.params.ProjectID); err != nil {
			return err
		}
	} else {
		opt := option.WithCredentialsFile(hlpr.params.CredPath)
		if hlpr.fsClient, err = firestore.NewClient(ctx, hlpr.params.ProjectID, opt); err != nil {
			return err
		}
	}
	return nil
}
func (hlpr *FsClientHelper) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	/*if param == nil {
		return fmt.Errorf("FsClientHelper.open: nil param")
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
			return fmt.Errorf("FsClientHelper.open: firebase projectID param not specified")
		}
		var ok bool
		hlpr.params.ProjectID, ok = pid.(string)
		if !ok {
			return fmt.Errorf("FsClientHelper.open: firebase projectID param not a string")
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
		} else {
			return fmt.Errorf("FsClientHelper.open: firebase credential param not specified")
		}
	}
	return nil
}

func (hlpr *FsClientHelper) close() {
	if hlpr.fsClient != nil {
		hlpr.fsClient.Close()
		hlpr.fsClient = nil
	}

}
