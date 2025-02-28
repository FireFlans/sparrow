package utils

import (
	"fmt"
	"sparrow/structures"
)

func ParseXMLLabel(label structures.OriginatorConfidentialityLabel) map[string]interface{} {
	response := map[string]interface{}{
		"PolicyIdentifier": label.ConfidentialityInformation.PolicyIdentifier,
		"Classification":   label.ConfidentialityInformation.Classification,
		"Categories":       map[string]map[string]interface{}{},
	}

	// Populate the categories
	for _, category := range label.ConfidentialityInformation.Categories {
		fmt.Println(category.TagName)
		response["Categories"].(map[string]map[string]interface{})[category.TagName] = map[string]interface{}{
			"type":   category.Type,
			"values": category.GenericValues,
		}
	}
	return response
}
