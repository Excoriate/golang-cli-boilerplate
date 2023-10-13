package utils

import "fmt"

func FindStrInSlice(slice []string, str string) (string, error) {
	for _, s := range slice {
		if s == str {
			return s, nil
		}
	}

	return "", fmt.Errorf("failed to find string in slice")
}

func FindStrInMapValues(m map[string]string, str string) (string, error) {
	for _, v := range m {
		if v == str {
			return v, nil
		}
	}

	return "", fmt.Errorf("failed to find string in map values")
}

func FindStrInMapKeys(m map[string]string, str string) (string, error) {
	for k := range m {
		if k == str {
			return k, nil
		}
	}

	return "", fmt.Errorf("failed to find string in map keys")
}

func KeyInMapHasValueSet(m map[string]string, keyInMap string) error {
	// If the keyInMap is empty, return an error.
	if keyInMap == "" {
		return fmt.Errorf("failed to find string in map, keyInMap is empty")
	}

	for k, v := range m {
		if k == keyInMap {
			if v == "" {
				return fmt.Errorf("failed to find string in map, keyInMap is empty")
			}
		}
	}

	return nil
}
