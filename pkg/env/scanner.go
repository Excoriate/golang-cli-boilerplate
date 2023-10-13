package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

// ScanEnvVarsFromHost CreateWebRoute all env vars.
func ScanEnvVarsFromHost() (types.EnvVars, error) {
	result := make(types.EnvVars)

	for _, env := range os.Environ() {
		keyValue := strings.Split(env, "=")
		result[keyValue[0]] = utils.RemoveDoubleQuotes(keyValue[1])
	}

	return result, nil
}

// ScanEnvVarsWithPrefix CreateWebRoute env vars with prefix.
// Returns an error if any of the variables either do not exist or have an empty value.
func ScanEnvVarsWithPrefix(prefix string) (types.EnvVars, error) {
	result := make(types.EnvVars)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key := pair[0]

		if strings.HasPrefix(key, prefix) {
			value := pair[1]
			if value == "" {
				return nil, fmt.Errorf("environment variable %s has an empty value", key)
			}
			result[key] = utils.RemoveDoubleQuotes(value)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no environment variables with the prefix %s found", prefix)
	}

	return result, nil
}
