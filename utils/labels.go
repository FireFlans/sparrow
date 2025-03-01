package utils

import (
	"encoding/xml"
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

func GenerateXMLLabel(jsonData structures.JSONConfidentialityLabel) string {
	var categories []structures.Category
	for tag, cat := range jsonData.Categories {
		categories = append(categories, structures.Category{
			TagName:       tag,
			Type:          cat.Type,
			GenericValues: cat.Values,
		})
	}
	xmlData := structures.OriginatorConfidentialityLabel{
		XMLNS: "urn:nato:stanag:4774:confidentialitymetadatalabel:1:0",
		ConfidentialityInformation: structures.ConfidentialityInformation{
			PolicyIdentifier: jsonData.PolicyIdentifier,
			Classification:   jsonData.Classification,
			Categories:       categories,
		},
	}

	output, err := xml.MarshalIndent(xmlData, "", "  ")
	if err != nil {
		return ""
	}

	return "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n" + string(output)

}
