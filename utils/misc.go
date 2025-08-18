package utils

import "strings"

func IntersectStringsArrays(a, b []string) []string {
	setA := make(map[string]bool)
	for _, item := range a {
		setA[item] = true
	}

	result := []string{}
	alreadyIn := make(map[string]bool)
	for _, item := range b {
		if setA[item] && !alreadyIn[item] {
			result = append(result, item)
			alreadyIn[item] = true
		}
	}
	return result
}

func UnionStringArray(a, b []string) []string {
	set := make(map[string]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	for _, v := range b {
		set[v] = struct{}{}
	}
	var result []string
	for v := range set {
		result = append(result, v)
	}
	return result
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
