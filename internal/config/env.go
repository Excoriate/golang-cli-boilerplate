package config

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/env"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

// GetEnvVarsFromHost GetEnvVarsCfg returns the environment variables configuration
func GetEnvVarsFromHost() (types.EnvVars, error) {
	return env.ScanEnvVarsFromHost()
}

func GetEnvVarsFromDotEnv(dotEnvFile string) (types.EnvVars, error) {
	if dotEnvFile == "" {
		return nil, errors.New("failed to get env vars from dot env file: file path is empty")
	}

	gitRepoDir, err := utils.FindGitRepoDir(3)
	if err != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file "+
			"cant find the root directory for the .git repository: %s", err.Error())
	}

	dotEnvFilePath := filepath.Join(gitRepoDir, dotEnvFile)

	if err := utils.FileExistAndItIsAFile(dotEnvFilePath); err != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file: %s", err.Error())
	}

	if err := utils.FileIsNotEmpty(dotEnvFilePath); err != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file: %s", err.Error())
	}

	return env.GetEnvVarsFromDotFile(env.DotEnvFileReadOpt{
		FilePath:       dotEnvFilePath,
		IgnoreComments: true,
	})
}

func AttachEnvVarsToConfig(envVarsCfg, envVars types.EnvVars) (types.EnvVars, error) {
	if err := env.SetEnvVars(envVars); err != nil {
		return nil, fmt.Errorf("failed to attach env vars to config: %s", err.Error())
	}

	return env.MergeEnvVars(envVarsCfg, envVars), nil
}
