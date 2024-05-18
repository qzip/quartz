/*
Copyright (c) 2020, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

*/

package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/util"
	"reflect"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// Global Config params keys (when ParamPubSub.ProjectID & CredPath) are absent
const (
	CfgParamGcpPubSubPfix = "gcpPubSub"
	CfgParamProjectID     = "projectID"
	CfgParamCredPath      = "credentialsPath"
)

// GcpPubSubClientHelper GcpPubSub client
type GcpPubSubClientHelper struct {
	psClient *pubsub.Client
	params   *ParamGcpPubSubClientHelper
}

// ParamGcpPubSubClientHelper configuration parameters
type ParamGcpPubSubClientHelper struct {
	ProjectID string `json:"projectID,omitempty"`
	// CredPath is optional, or can be "default" or <path to cred file>
	CredPath string `json:"credentialsPath,omitempty"`
	//TimeDurationSec timer sleep interval while checking for done context
	TimeDurationSec time.Duration `json:"timeDurationSec,omitempty"`
}

// Name implements commands.Components Methods
func (hlpr *GcpPubSubClientHelper) Name() string {
	return "GcpPubSub.Client"
}

// Help implements commands.Components Methods
func (hlpr *GcpPubSubClientHelper) Help() string {
	return `
	 Helper for GcpPubSub client: all parameters, except projectID are optional
	 "params": {
		"projectID" : "my-project-id",
		"credPath" : "optional/path/to/file.json",
		"timeDurationSec" : 1,
	 }
	`
}

// ComponentType implements commands.Components Methods
func (hlpr *GcpPubSubClientHelper) ComponentType() reflect.Type {
	return reflect.TypeOf(hlpr)
}

// CreateHelper implements commands.HelperFactory Methods
func (hlpr *GcpPubSubClientHelper) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (psClient interface{}, err error) {
	if hlpr.psClient != nil {
		psClient = hlpr.psClient
		return
	}
	err = hlpr.open(ctx, param, cfg)
	if err == nil {
		var sleep = time.Duration(1)
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
	psClient = hlpr.psClient
	return
}

func (hlpr *GcpPubSubClientHelper) open(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	err := hlpr.getParams(ctx, param, cfg)
	if err != nil {
		return err
	}
	if hlpr.params.CredPath == "" || strings.EqualFold(hlpr.params.CredPath, "default") {
		if hlpr.psClient, err = pubsub.NewClient(ctx, hlpr.params.ProjectID); err != nil {
			return err
		}
	} else {
		opt := option.WithCredentialsFile(hlpr.params.CredPath)
		if hlpr.psClient, err = pubsub.NewClient(ctx, hlpr.params.ProjectID, opt); err != nil {
			return err
		}
	}
	return nil
}
func (hlpr *GcpPubSubClientHelper) getParams(_ context.Context, param interface{}, cfg map[string]interface{}) error {
	/*if param == nil {
		return fmt.Errorf("GcpPubSubClientHelper.open: nil param")
	}*/
	hlpr.params = &ParamGcpPubSubClientHelper{}
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
		pid := util.FlattenMap(cfg, CfgParamGcpPubSubPfix+"."+CfgParamProjectID)
		if pid == nil {
			return fmt.Errorf("GcpPubSubClientHelper.open: GcpPubSub projectID param not specified")
		}
		var ok bool
		hlpr.params.ProjectID, ok = pid.(string)
		if !ok {
			return fmt.Errorf("GcpPubSubClientHelper.open: GcpPubSub projectID param not a string")
		}
	}
	// check if credential path present
	if hlpr.params.CredPath == "" {
		pid := util.FlattenMap(cfg, CfgParamGcpPubSubPfix+"."+CfgParamCredPath)
		if pid == nil {
			// assumed default
			hlpr.params.CredPath = ""

		}
		if p, ok := pid.(string); ok {
			hlpr.params.CredPath = p
		} else {
			return fmt.Errorf("GcpPubSubClientHelper.open: GcpPubSub credential param not specified")
		}
	}
	return nil
}

func (hlpr *GcpPubSubClientHelper) close() {
	if hlpr.psClient != nil {
		hlpr.psClient.Close()
		hlpr.psClient = nil
	}

}
