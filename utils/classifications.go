package utils

import (
	"errors"
	"sparrow/structures"
)

func GetClassifications(spifs []structures.SPIF, policy string) []string {
	var classifications []string
	desiredSpif, err := FindPolicy(spifs, policy)
	if err != nil {
		return classifications
	}
	for _, classification := range desiredSpif.SecurityClassifications.Classifications {
		classifications = append(classifications, classification.Name)
	}

	return classifications
}

func GetClassificationHierarchy(spifs []structures.SPIF, policyIdentifier string, classification string) int {
	var spif, err = FindPolicy(spifs, policyIdentifier)
	if err != nil {
		return -1
	}
	for _, spifClassification := range spif.SecurityClassifications.Classifications {
		if spifClassification.Name == classification {
			return spifClassification.Hierarchy
		}
	}
	return -1
}

func GetHierarchyClassification(spifs []structures.SPIF, policyIdentifier string, hierarchy int) (string, error) {
	var spif, err = FindPolicy(spifs, policyIdentifier)
	if err != nil {
		return "", errors.New("cannot find required policy")
	}
	for _, spifClassification := range spif.SecurityClassifications.Classifications {
		if spifClassification.Hierarchy == hierarchy {
			return spifClassification.Name, nil
		}
	}
	return "", errors.New("cannot find required classification")
}
