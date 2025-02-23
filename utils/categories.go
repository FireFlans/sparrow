package utils

func GetCategories(spifs []SPIF, policy string, classification string) []string {
	var categories []string
	desiredSpif, err := FindPolicy(spifs, policy)
	if err != nil {
		return categories
	}
	for _, tagSet := range desiredSpif.SecurityCategoryTagSets.TagSet {

		for _, securityTag := range tagSet.Tag {
			categories = append(categories, securityTag.Name)
		}
	}
	return categories
}

func GetType(spifs []SPIF, policy string, category string) (response string) {
	desiredSpif, err := FindPolicy(spifs, policy)
	if err != nil {
		return ""
	}

	for _, tagSet := range desiredSpif.SecurityCategoryTagSets.TagSet {
		for _, tag := range tagSet.Tag {
			if tag.Name == category {
				return tag.TagType
			}
		}
	}
	return ""
}
