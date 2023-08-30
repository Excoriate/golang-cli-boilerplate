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
