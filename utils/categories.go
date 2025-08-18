package utils

import (
	"fmt"
	"sparrow/structures"
)

func GetCategories(spifs []structures.SPIF, policy string, classification string) []string {
	var categories []string
	desiredSpif, err := FindPolicy(spifs, policy)
	if err != nil {
		return categories
	}
	for _, tagSet := range desiredSpif.SecurityCategoryTagSets.TagSet {
		var NoMentions = GetMentions(spifs, policy, classification, tagSet.SecurityCategoryTag.Name)
		if len(NoMentions) == 0 {
			continue //Discard the Category as no values are in it
		}
		categories = append(categories, tagSet.SecurityCategoryTag.Name)

	}
	return categories
}

func GetType(spifs []structures.SPIF, policy string, category string) (response string) {
	desiredSpif, err := FindPolicy(spifs, policy)
	if err != nil {
		return ""
	}
	for _, tagSet := range desiredSpif.SecurityCategoryTagSets.TagSet {

		if tagSet.SecurityCategoryTag.Name == category {
			if tagSet.SecurityCategoryTag.TagType == "enumerated" {
				return tagSet.SecurityCategoryTag.EnumType
			}
			return tagSet.SecurityCategoryTag.TagType
		}

	}
	return ""
}

func ExtractPermissiveCategories(label structures.JSONConfidentialityLabel) map[string]string {
	permissiveCategories := make(map[string]string)
	for _, category := range label.Categories {

		if category.Type == "PERMISSIVE" {
			fmt.Println(category)
		}
	}
	return permissiveCategories
}

func GetRestrictiveCategories(label structures.JSONConfidentialityLabel) []string {
	var restrictiveCategories []string
	for _, category := range label.Categories {
		if category.Type == "RESTRICTIVE" {
			restrictiveCategories = append(restrictiveCategories, category.Type)
		}
	}
	return restrictiveCategories
}

func FindCategoryIndex(categories []structures.Category, newCategory string) int {
	for index, category := range categories {
		if category.TagName == newCategory {
			return index
		}
	}
	return -1
}
