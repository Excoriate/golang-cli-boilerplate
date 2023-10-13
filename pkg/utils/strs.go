package utils

import (
	"strings"

	"github.com/google/uuid"
)

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

// WrapAtLength wraps the input paragraph at the specified length
func WrapAtLength(input string, length int) string {
	words := strings.Fields(input)
	if len(words) == 0 {
		return input
	}

	wrapped := ""
	line := words[0]
	for _, word := range words[1:] {
		if len(line)+1+len(word) <= length {
			line += " " + word
		} else {
			wrapped += "    " + line + "\n"
			line = word
		}
	}
	// Add the last line
	wrapped += "    " + line
	return wrapped
}

func CleanPathFromExtraSlashes(target string) string {
	return strings.ReplaceAll(target, "//", "/")
}

func GetUUID() string {
	return uuid.New().String()
}
