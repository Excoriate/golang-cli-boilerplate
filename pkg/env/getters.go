package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type VarSkeleton struct {
	Key   string
	Value string
}

func GetFromMultiplePotentialKeys(keysAllowed []string) (VarSkeleton, error) {
	if len(keysAllowed) == 0 {
		return VarSkeleton{}, errors.New("no keys allowed")
	}

	for _, key := range keysAllowed {
		value, isSet := os.LookupEnv(key)
		if isSet && value != "" {
			return VarSkeleton{
				Key:   key,
				Value: value,
			}, nil
		}
	}

	return VarSkeleton{}, fmt.Errorf("no keys allowed were set. Keys allowed: %s", strings.Join(keysAllowed, ", "))
}

// GetStringOrDefault CreateWebRoute string value from Viper by given key.
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

// GetNumberOrDefault CreateWebRoute number value from Viper by given key.
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

func GetNumberInt64OrDefault(key string, defaultValue int64) int64 {
	if key == "" {
		return defaultValue
	}

	value, isSet := os.LookupEnv(key)

	if !isSet {
		return defaultValue
	}

	// Parse value to int
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

// GetBoolOrDefault CreateWebRoute boolean value from Viper by given key.
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

// GetDurationOrDefault CreateWebRoute duration value from Viper by given key.
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

func GetSliceOrDefault(key string, defaultValue []string) []string {
	if key == "" {
		return defaultValue
	}

	value, isSet := os.LookupEnv(key)
	if !isSet {
		return defaultValue
	}

	return strings.Split(value, ",")
}

func IsSet(key string) error {
	_, isSet := os.LookupEnv(key)
	if !isSet {
		return fmt.Errorf("env var %s is not set", key)
	}

	return nil
}
