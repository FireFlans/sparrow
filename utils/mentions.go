package utils

import (
	"sparrow/structures"
)

/*
Returns all values from a security category
*/
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
			if !Contains(tagCategory.ExcludedClasses, classification) {
				mentions = append(mentions, tagCategory.Name)
			}
		}
	}
	return mentions
}

/*
Returns values from a mention in a label
*/
