package env

import (
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
)

func ScanTerraformEnvVarsFromHost() (types.EnvVars, error) {
	envVarsWithPrefix, err := ScanEnvVarsWithPrefix("TF_VAR_")
	if err != nil {
		return nil, fmt.Errorf("failed to scan env vars with prefix TF_VAR_: %w", err)
	}

	return envVarsWithPrefix, nil
}
