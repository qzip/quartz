package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/util"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// Global Config params keys (when ParamPubSub.ProjectID & CredPath) are absent
const (
	CfgParamGcpPubSubPfix = "gcpPubSub"
	CfgParamProjectID     = "projectID"
	CfgParamCredPath      = "credentialsPath"
)

//MsgConsumerHelper is created by the helper & for each message in queue, processes Messages sequentially,
//if error is nil, then ACK is sent to the message queue, return error.
type MsgConsumerHelper interface {
	Open(ctx context.Context) error
	Consume(ctx context.Context, m *pubsub.Message) error
	Close(ctx context.Context)
}

//CfgProcessQueue load the parameters.
type CfgProcessQueue struct {
	ProjectID string `json:"projectID,omitempty"`
	// CredPath is optional, or can be "default" or <path to cred file>
	CredPath     string `json:"credentialsPath,omitempty"`
	Subscription string `json:"subscription"`
	//MsgConsumerHelper  is the message processor helper
	MsgConsumerHelper string `json:"messageConsumerHelper"`
}

//ProcessGcpQueue processes the GCP PubSub Subscription Queue
type ProcessGcpQueue struct {
	errChan      chan error
	psClient     *pubsub.Client
	msgProcessor MsgConsumerHelper
	params       *CfgProcessQueue
}

// Name implements command interface method
func (pq *ProcessGcpQueue) Name() string {
	return "cmd.ProcessGcpQueue"
}

//Help implements command interface method
func (pq *ProcessGcpQueue) Help() string {
	return `
	  # cmd.ProcessGcpQueue: processes GCP Subscriptions
	  {
		  "command" : "cmd.ProcessGcpQueue",
		  "param" : {
			
			"messageProcessor": {
				"component": "helper.Msg2File",
				"param": {
					#component specific parameters
					#TODO
					"messageConsumerHelper": "<helper component Name>"
				}
			}

			}
	  }
	`
}

/*
commands.CfgHelpersKey

*/

//Exec executes the stages, if the prev. stage is successful
func (pq *ProcessGcpQueue) Exec(ctx context.Context, cfg map[string]interface{}, errChan chan error) {
	if err := pq.open(ctx, cfg["param"], cfg); err != nil {
		errChan <- err
		return
	}
	defer pq.close()
	if err := pq.loadMsgProc(ctx); err != nil {
		errChan <- err
		return
	}
	if err := pq.msgProcessor.Open(ctx); err != nil {
		errChan <- err
		return
	}
	defer pq.msgProcessor.Close(ctx)
	subs := pq.psClient.Subscription(pq.params.Subscription)
	if ok, err := subs.Exists(ctx); !ok || err != nil {
		if err != nil {
			errChan <- err
			return
		}
		if !ok {
			errChan <- commands.NewFatalError(fmt.Sprintf("ProcessGcpQueue.Exec: %v subscription does not exist", pq.params.Subscription))
			return
		}
	}

	cctx, cancel := context.WithCancel(ctx)
	processor := func(ctx context.Context, m *pubsub.Message) {
		select {
		case <-ctx.Done():
			util.DebugInfo(ctx, "ProcessGcpQueue.Exec: ctx done\n\n")
			m.Nack()
			return // for
		default:
		}
		if err := pq.msgProcessor.Consume(ctx, m); err != nil {
			errChan <- err
			m.Nack()
			cancel()
		} else {
			m.Ack()
		}
	}

	errChan <- subs.Receive(cctx, processor)
}

func (pq *ProcessGcpQueue) loadMsgProc(ctx context.Context) error {
	helper := ctx.Value(pq.params.MsgConsumerHelper)
	if helper == nil {
		return commands.NewFatalError(fmt.Sprintf("ProcessGcpQueue.loadMsgProc: %v message processor component not registered", pq.params.MsgConsumerHelper))
	}
	ok := false
	pq.msgProcessor, ok = helper.(MsgConsumerHelper)
	if !ok {
		return commands.NewFatalError(fmt.Sprintf("ProcessGcpQueue.loadMsgProc: %v component not of MsgConsumerHelper type", pq.params.MsgConsumerHelper))
	}
	return nil
}
func (pq *ProcessGcpQueue) open(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	err := pq.getParams(ctx, param, cfg)
	if err != nil {
		return err
	}
	if pq.params.CredPath == "" || strings.EqualFold(pq.params.CredPath, "default") {
		if pq.psClient, err = pubsub.NewClient(ctx, pq.params.ProjectID); err != nil {
			return err
		}
	} else {
		opt := option.WithCredentialsFile(pq.params.CredPath)
		if pq.psClient, err = pubsub.NewClient(ctx, pq.params.ProjectID, opt); err != nil {
			return err
		}
	}
	return nil
}
func (pq *ProcessGcpQueue) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	by, err := json.Marshal(param)
	if err != nil {
		return err
	}
	cfgQue := &CfgProcessQueue{}
	if err := json.Unmarshal(by, cfgQue); err != nil {
		return err
	}
	if pq.params.ProjectID == "" {
		pid := util.FlattenMap(cfg, CfgParamGcpPubSubPfix+"."+CfgParamProjectID)
		if pid == nil {
			return fmt.Errorf("ProcessGcpQueue.open: GcpPubSub projectID param not specified")
		}
		var ok bool
		pq.params.ProjectID, ok = pid.(string)
		if !ok {
			return fmt.Errorf("ProcessGcpQueue.open: GcpPubSub projectID param not a string")
		}
	}
	// check if credential path present
	if pq.params.CredPath == "" {
		pid := util.FlattenMap(cfg, CfgParamGcpPubSubPfix+"."+CfgParamCredPath)
		if pid == nil {
			// assumed default
			pq.params.CredPath = ""

		}
		if p, ok := pid.(string); ok {
			pq.params.CredPath = p
		} else {
			return fmt.Errorf("ProcessGcpQueue.open: GcpPubSub credential param not specified")
		}
	}
	return nil
}
func (pq *ProcessGcpQueue) close() {
	if pq.psClient != nil {
		pq.psClient.Close()
		pq.psClient = nil
	}

}
