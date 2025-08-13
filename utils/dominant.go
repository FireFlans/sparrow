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

	// Permissive categories related stuff [Rules 4-5-6]

	ExtractPermissiveCategories(labels[0])
	/*
		for _, label := range labels[1:] {
			// Get current label permissive categories
			labelPermissiveCategories := GetPermissiveCategories(label)

			// Intersect with previous categories as we only want categories present in all labels [Rules - 6]
			permissiveCategories = IntersectStringsArrays(permissiveCategories, labelPermissiveCategories)
		}*/

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
