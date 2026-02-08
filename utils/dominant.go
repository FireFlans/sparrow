package utils

import (
	"errors"
	"sparrow/structures"
)

func DominantLabel(spifs []structures.SPIF, labels []structures.JSONConfidentialityLabel) (structures.OriginatorConfidentialityLabel, error) {
	dominantLabel := structures.OriginatorConfidentialityLabel{
		XMLNS: "urn:nato:stanag:4774:confidentialitymetadatalabel:1:0",
		ConfidentialityInformation: structures.ConfidentialityInformation{
			PolicyIdentifier: "",
			Classification:   "",
			Categories:       nil,
		},
	}
	//[Rule - 1]
	policy, err := dominantRule1(labels)
	if err != nil {
		/*[Rule - 2]
		In  the  case  that  a  confidentiality  label  cannot  be  mapped  to  an
		equivalent governing policy confidentiality label (for whatever reason)
		a policy decision is required. This may be that human intervention is
		required or the default system high confidentiality label is applied.
		*/
		return dominantLabel, err
	}
	dominantLabel.ConfidentialityInformation.PolicyIdentifier = policy

	classification, err := dominantRule3(spifs, labels, policy)
	if err != nil {
		return dominantLabel, err
	}
	dominantLabel.ConfidentialityInformation.Classification = classification

	// Security categories related stuff [Rules 4-5-6-7-8]

	// Group all categories by type
	categoryMap := make(map[string][]structures.Category)
	for _, label := range labels {
		for tag, jsonCat := range label.Categories {
			cat := structures.Category{
				TagName:       tag,
				Type:          jsonCat.Type,
				GenericValues: jsonCat.Values,
			}
			categoryMap[cat.Type] = append(categoryMap[cat.Type], cat)
		}
	}
	// Build the final list of security categories
	var securityCategories []structures.Category
	for catType, cats := range categoryMap {
		// Rule 4-5-6: Permissive categories
		if catType == "PERMISSIVE" {
			/*
				[Rule - 6] If a confidentiality label (within a set of confidentiality labels) does not
				contain  a  permissive  category  (of  the  same  type)  that  one  or  more
				confidentiality  labels  contain  the  dominant  confidentiality  label  must
				not contain that permissive category.
			*/
			allLabelsHaveCategory := true
			for _, label := range labels {
				if _, ok := label.Categories[cats[0].TagName]; !ok {
					allLabelsHaveCategory = false
					break
				}
			}
			if !allLabelsHaveCategory {
				continue
			}
			/*
				[Rule - 4] For all confidentiality labels that contain a permissive category (of the
				same type) the dominant confidentiality label must contain a
				permissive category (of that type) with the intersection of the category
				values.
			*/
			var intersection []string
			if len(cats) > 0 {
				intersection = cats[0].GenericValues
				for _, cat := range cats[1:] {
					intersection = IntersectStringsArrays(intersection, cat.GenericValues)
				}
			}
			/*
				[Rule - 5] For all confidentiality labels that contain a permissive category (of the
				same type) and the intersection of the category values is empty the
				dominant confidentiality label must not contain that permissive
				category.
			*/
			if len(intersection) > 0 {
				securityCategories = append(securityCategories, structures.Category{
					TagName:       cats[0].TagName,
					Type:          catType,
					GenericValues: intersection,
				})
			}
		} else if catType == "RESTRICTIVE" {
			/*
				[Rule - 7] For each confidentiality label that exists with one or more restrictive
				categories (of the same type) the dominant confidentiality label must
				contain  a  restrictive  category  (of  that  type)  with  the  union  of  the
				category values.
			*/
			var union []string
			for _, cat := range cats {
				union = UnionStringArray(union, cat.GenericValues)
			}
			securityCategories = append(securityCategories, structures.Category{
				TagName:       cats[0].TagName,
				Type:          catType,
				GenericValues: union,
			})
		} else {
			/*
				[Rule - 8] For each confidentiality label that exists with one or more informative
				categories (of the same type) the dominant confidentiality label may
				contain  an  informative  category  (of  that  type)  with  the  union  of  the
				category values.
				INCLUDING THEM IN THE DOMINANT LABEL
			*/
			var union []string
			for _, cat := range cats {
				union = UnionStringArray(union, cat.GenericValues)
			}
			securityCategories = append(securityCategories, structures.Category{
				TagName:       cats[0].TagName,
				Type:          catType,
				GenericValues: union,
			})
		}
	}

	dominantLabel.ConfidentialityInformation.Categories = securityCategories

	return dominantLabel, nil
}

/*
[Rule - 1] All  confidentiality  labels  must  have  the  same  policy  identifier  (the
governing organization policy that is being enforced).
[TODO]
Any confidentiality label containing a foreign policy must be mapped to the
equivalent governing policy and stored as an alternative
confidentiality label.
*/
func dominantRule1(labels []structures.JSONConfidentialityLabel) (string, error) {
	policy := labels[0].PolicyIdentifier
	for _, label := range labels[1:] {
		if policy != label.PolicyIdentifier {
			return "", errors.New("cannot determine a dominant policy")
		}
	}
	return policy, nil
}

/*
[Rule - 3] The dominant confidentiality label classification value must be
determined based on the classification hierarchy value as specified in
the XML SPIF.
*/

func dominantRule3(spifs []structures.SPIF, labels []structures.JSONConfidentialityLabel, policy string) (string, error) {

	classificationHierarchy := GetClassificationHierarchy(spifs, labels[0].PolicyIdentifier, labels[0].Classification)
	for _, label := range labels[1:] {
		labelClassificationHierarchy := GetClassificationHierarchy(spifs, label.PolicyIdentifier, label.Classification)
		if classificationHierarchy < labelClassificationHierarchy {
			classificationHierarchy = labelClassificationHierarchy
		}
	}
	classification, err := GetHierarchyClassification(spifs, policy, classificationHierarchy)
	if err != nil {
		return "", err
	}
	return classification, nil

}
