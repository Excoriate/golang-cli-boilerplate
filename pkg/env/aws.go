package env

import (
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
)

func ScanAWSEnvVarsFromHost() (types.EnvVars, error) {
	envVarsWithPrefix, err := ScanEnvVarsWithPrefix("AWS_")
	if err != nil {
		return nil, fmt.Errorf("failed to scan env vars with prefix AWS_: %w", err)
	}

	return envVarsWithPrefix, nil
}
