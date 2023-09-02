package utils

import "strings"

func RemoveDoubleQuotes(target string) string {
	return strings.Trim(target, "\"")
}

func NormaliseStringToLower(target string) string {
	if target == "" {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(target))
}

func NormaliseStringToUpper(target string) string {
	if target == "" {
		return ""
	}
	return strings.ToUpper(strings.TrimSpace(target))
}

func MergeSlices(slices ...[]string) []string {
	var merged []string
	for _, slice := range slices {
		merged = append(merged, slice...)
	}
	return merged
}
