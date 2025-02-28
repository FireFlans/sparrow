package structures

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
	PrivilegeId             string                  `xml:"privilegeId,attr"`
	RbacId                  string                  `xml:"rbacId,attr"`
	SecurityPolicyId        SecurityPolicyID        `xml:"securityPolicyId"`
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
	Name             string           `xml:"name,attr"`
	Lacv             int              `xml:"lacv,attr"`
	Hierarchy        int              `xml:"hierarchy,attr"`
	MarkingData      MarkingData      `xml:"markingData"`
	MarkingQualifier MarkingQualifier `xml:"markingQualifier"`
}

type MarkingQualifier struct {
	MarkingCode string      `xml:"markingCode,attr"`
	Qualifiers  []Qualifier `xml:"qualifier"`
}

type MarkingData struct {
	Lang   string `xml:"lang,attr"`
	Phrase string `xml:"phrase,attr"`
	Code   string `xml:"code"`
}

type RequiredCategory struct {
	Operation string `xml:"operation,attr"`
}

type Qualifier struct {
	MarkingQualifier string `xml:"markingQualifier,attr"`
	QualifierCode    string `xml:"qualifierCode,attr"`
}

type SecurityCategoryTagSets struct {
	TagSet []SecurityCategoryTagSet `xml:"securityCategoryTagSet"`
}

type SecurityCategoryTagSet struct {
	Name                string              `xml:"name,attr"`
	ID                  string              `xml:"id,attr"`
	SecurityCategoryTag SecurityCategoryTag `xml:"securityCategoryTag"`
}

type SecurityCategoryTag struct {
	Name            string        `xml:"name,attr"`
	TagType         string        `xml:"tagType,attr"`
	SingleSelection bool          `xml:"singleSelection,attr"`
	TagCategories   []tagCategory `xml:"tagCategory"`
}

type tagCategory struct {
	Name            string   `xml:"name,attr"`
	LACV            string   `xml:"lacv,attr"`
	ExcludedClasses []string `xml:"excludedClass"`
	Obsolete        bool     `xml:"obsolete,attr"`
}
