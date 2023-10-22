/*
Copyright (c) 2023, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package ondc

type GroupElement struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	ParentID    string                 `json:"parentid,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"` //needs to have a metadata schema disctionary
}
