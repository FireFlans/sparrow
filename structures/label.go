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
	XMLNS                      string                     `xml:"xmlns:s4774,attr"`
	ConfidentialityInformation ConfidentialityInformation `xml:"ConfidentialityInformation"`
}

type JSONCategory struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type JSONConfidentialityLabel struct {
	PolicyIdentifier string                  `json:"PolicyIdentifier"`
	Classification   string                  `json:"Classification"`
	Categories       map[string]JSONCategory `json:"Categories"`
}
