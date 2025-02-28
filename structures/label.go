package structures

import (
	"encoding/xml"
)

type ConfidentialityInformation struct {
	PolicyIdentifier string     `xml:"PolicyIdentifier"`
	Classification   string     `xml:"Classification"`
	Categories       []Category `xml:"Category"`
}
type Category struct {
	TagName       string   `xml:"TagName,attr"`
	Type          string   `xml:"Type,attr"`
	GenericValues []string `xml:"GenericValue"`
}

type OriginatorConfidentialityLabel struct {
	XMLName                    xml.Name                   `xml:"originatorConfidentialityLabel"`
	ConfidentialityInformation ConfidentialityInformation `xml:"ConfidentialityInformation"`
}
