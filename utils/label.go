package utils

import (
	"encoding/xml"
)

/*
Dummy Label struct, TODO
*/
type Label struct {
	XMLName       xml.Name `xml:"SPIF"`
	SchemaVersion string   `xml:"schemaVersion,attr"`
	Version       string   `xml:"version,attr"`
	CreationDate  string   `xml:"creationDate,attr"`
	OriginatorDN  string   `xml:"originatorDN,attr"`
	KeyIdentifier string   `xml:"keyIdentifier,attr"`
	PrivilegeID   string   `xml:"privilegeId,attr"`
	RBACID        string   `xml:"rbacId,attr"`
}
