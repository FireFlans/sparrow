package utils

import (
	"encoding/xml"
)

type SPIF struct {
	XMLName                 xml.Name                `xml:"SPIF"`
	SchemaVersion           string                  `xml:"schemaVersion,attr"`
	Version                 string                  `xml:"version,attr"`
	CreationDate            string                  `xml:"creationDate,attr"`
	OriginatorDN            string                  `xml:"originatorDN,attr"`
	KeyIdentifier           string                  `xml:"keyIdentifier,attr"`
	PrivilegeID             string                  `xml:"privilegeId,attr"`
	RBACID                  string                  `xml:"rbacId,attr"`
	SecurityPolicyID        SecurityPolicyID        `xml:"securityPolicyId"`
	SecurityClassifications SecurityClassifications `xml:"securityClassifications"`
	SecurityCategoryTagSets SecurityCategoryTagSets `xml:"securityCategoryTagSets"`
}

type SecurityPolicyID struct {
	Name string `xml:"name,attr"`
	ID   string `xml:"id,attr"`
}

type SecurityClassifications struct {
	Classifications []SecurityClassification `xml:"securityClassification"`
}

type SecurityClassification struct {
	Name      			string `xml:"name,attr"`
	LACV      			string `xml:"lacv,attr"`
	Hierarchy 			string `xml:"hierarchy,attr"`
	Color	  			string `xml:"color,attr"`
	RequiredCategory	RequiredCategory	`xml:"requiredCategory"`
}

type RequiredCategory struct {
	Operation	string `xml:"operation,attr"`
}

type SecurityCategoryTagSets struct {
	TagSet []SecurityCategoryTagSet `xml:"securityCategoryTagSet"`
}

type SecurityCategoryTagSet struct {
	Name string                `xml:"name,attr"`
	ID   string                `xml:"id,attr"`
	Tag  []SecurityCategoryTag `xml:"securityCategoryTag"`
}

type SecurityCategoryTag struct {
	Name            string        `xml:"name,attr"`
	TagType         string        `xml:"tagType,attr"`
	SingleSelection bool          `xml:"singleSelection,attr"`
	Category        []tagCategory `xml:"tagCategory"`
}

type tagCategory struct {
	Name            string          `xml:"name,attr"`
	LACV            string          `xml:"lacv,attr"`
	ExcludedClasses []ExcludedClass `xml:"excludedClass"`
	Obsolete        bool            `xml:"obsolete,attr"`
}

type ExcludedClass struct {
	Value string `xml:",chardata"`
}
