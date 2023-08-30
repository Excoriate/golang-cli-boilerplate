package env

import (
	"errors"
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
)

func SetEnvVars(envVars types.EnvVars) error {
	if len(envVars) == 0 {
		return errors.New("failed to set env vars: env vars are empty")
	}

	for k, v := range envVars {
		if _, exists := os.LookupEnv(k); !exists {
			// If the environment variable doesn't exist, set it.
			if err := os.Setenv(k, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func MergeEnvVars(envVars ...map[string]string) map[string]string {
	result := make(map[string]string)

	for _, env := range envVars {
		for key, value := range env {
			if key != "" && value != "" {
				result[key] = utils.RemoveDoubleQuotes(value)
			}
		}
	}

	return result
}
