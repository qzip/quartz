/*
Copyright (c) 2021, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package firebase

// FbCollectionSaver implements seq.RunSeq
// This component may be called with different arguments
// by difference factory classes, so should not store
// instance info.
type FbCollectionSaver struct {
}

// CfgFbCollectionSaver contructs init params
type CfgFbCollectionSaver struct {
	FbClientCtxKey string `json:"FbClientCtxKey,omitempty"`
	CollectionName string `json:"collectionName"`
	DataInCtxName  string `json:"dataInCtxName"`
}
