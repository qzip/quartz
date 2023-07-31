/*
Copyright (c) 2020, QzIP Blockchain Technology LLP

All Rights Reserved

Author: Ashish Banerjee, <ashish@qzip.in>

*/

package dlt

/*
DeploymentDescriptor contains the genesis (per contract)
deployment json + template are added in a zip
& distributed to all the participants.
*/
type DeploymentDescriptor struct {
	InstanceID   string       `xml:"instanceID" json:"instanceID"`
	TemplateURL  string       `xml:"templateURL" json:"templateURL"`
	TemplateHash []byte       `xml:"templateHash" json:"templateHash"`
	Narration    Triple       `xml:"narration" json:"narration"`
	Documents    []OpaqueData `xml:"documents" json:"documents"`
	Partners     []Partner    `xml:"partners" json:"partners"`
	Tags         []Triple     `xml:"tags" json:"tags"`
}

//Partner partner details
type Partner struct {
	PartnerName        string
	PartnerDomain      string
	PartnerRoles       []string
	Email              string
	SignKeyFingerprint Triple
	// Block Exchange Protocol v1
	EndpointURI string   // eg., "net.syncthing:<syncthing address>"
	Tags        []Triple `xml:"tags" json:"tags"`
}
