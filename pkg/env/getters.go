package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
)

func GetFromMultiplePotentialKeys(keysAllowed []string) (types.EnvVar, error) {
	if len(keysAllowed) == 0 {
		return types.EnvVar{}, errors.New("no keys allowed")
	}

	for _, key := range keysAllowed {
		value, isSet := os.LookupEnv(key)
		if isSet && value != "" {
			return types.EnvVar{
				Key:   key,
				Value: value,
			}, nil
		}
	}

	return types.EnvVar{}, fmt.Errorf("no keys allowed were set. Keys allowed: %s", strings.Join(keysAllowed, ", "))
}

// GetStringOrDefault Get string value from Viper by given key.
func GetStringOrDefault(key, defaultValue string) string {
	if key == "" {
		return defaultValue
	}

	value, isSet := os.LookupEnv(key)
	if !isSet {
		return defaultValue
	}

	return value
}

// GetNumberOrDefault Get number value from Viper by given key.
func GetNumberOrDefault(key string, defaultValue int) int {
	if key == "" {
		return defaultValue
	}

	value, isSet := os.LookupEnv(key)

	if !isSet {
		return defaultValue
	}

	// Parse value to int
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

// GetBoolOrDefault Get boolean value from Viper by given key.
func GetBoolOrDefault(key string, defaultValue bool) bool {
	if key == "" {
		return defaultValue
	}

	value, isSet := os.LookupEnv(key)
	if !isSet {
		return defaultValue
	}

	// Parse value to bool
	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

// GetDurationOrDefault Get duration value from Viper by given key.
func GetDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if key == "" {
		return defaultValue
	}

	value, isSet := os.LookupEnv(key)
	if !isSet {
		return defaultValue
	}

	// Parse value to duration
	parsedValue, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}
