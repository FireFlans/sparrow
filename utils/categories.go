package utils

import (
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
			return tagSet.SecurityCategoryTag.TagType
		}

	}
	return ""
}
