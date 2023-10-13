package utils

import "fmt"

func MapIsNulOrEmpty(target map[string]string) bool {
	return len(target) == 0
}

func ValueInMapIsNotEmptyByKey(target map[string]string, key string) (string, error) {
	if MapIsNulOrEmpty(target) {
		return "", fmt.Errorf("failed to find string in map values")
	}

	for k, v := range target {
		if k == key {
			return v, nil
		}
	}

	return "", fmt.Errorf("failed to find string in map values")
}
