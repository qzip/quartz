/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

*/

/*
Package defs provides a configurable framework for commands
while avoiding multiple main packages.
*/
package defs

// Configuration parameters key values
const (
	SourcePipelinePrefix = "pipelines.source"
	ComponentSubKey      = "component"
	ParamSubKey          = "param"
	// array of sink
	SinkPipelinePrefix = "pipelines.sink"
	RtmDocSubColsChan  = "rtm$dschan"
	RunSeqParam        = "runners"
)

// Default helper context references
const (
	DefaFsClient  = "firestore.Client"
	DefaGcLogger  = "logging.Client"
	DefaBqClient  = "bigquery.Client"
	DefaLogName   = "ArtcilesLog"
	DefaBatchSize = 50
)

// Global Config params keys (when ParamDocSubColSrc.ProjectID & CredPath) are absent
const (
	CfgParamProjectID    = "projectID"
	CfgParamCredPath     = "credentialsPath"
	CfgParamFirebasePfix = "firebase"
)

// Context Helper keys
const (
	DocRecID = "docRecId"
	DocRec   = "docRec"
)

// BlockClassName default name for blockchain class identifier in serialization
const BlockClassName = "in.qzip.gpay.block"
