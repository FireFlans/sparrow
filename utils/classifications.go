package utils

import (
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
