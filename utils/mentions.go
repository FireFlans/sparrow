package utils

import (
	"sparrow/structures"
	"strings"
)

func GetMentions(spifs []structures.SPIF, policy string, classification string, category string) []string {
	var mentions []string
	desiredSpif, err := FindPolicy(spifs, policy)
	if err != nil {
		return mentions
	}
	for _, tagSet := range desiredSpif.SecurityCategoryTagSets.TagSet {
		// Check if the SecurityCategoryTag name matches (if provided)
		if category != "" && tagSet.SecurityCategoryTag.Name != category {
			continue
		}
		// Iterate through each TagCategory in the SecurityCategoryTag
		for _, tagCategory := range tagSet.SecurityCategoryTag.TagCategories {
			// Check if the classification is not in the excluded classes
			if !contains(tagCategory.ExcludedClasses, classification) {
				mentions = append(mentions, tagCategory.Name)
			}
		}
	}
	return mentions
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
