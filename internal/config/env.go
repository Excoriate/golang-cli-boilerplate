package config

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/env"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

type EnvVarManager struct {
}

type EnvVarFetcher interface {
	ScanEnvVarsFromHost() (types.EnvVars, error)
	ScanDotEnvAndAttachEnvVars(dotEnvFile string, currentEnvVars types.EnvVars) (types.EnvVars, error)
	AttachEnvVarsToConfig(envVarsCfg, envVars types.EnvVars) (types.EnvVars, error)
}

func (e *EnvVarManager) ScanEnvVarsFromHost() (types.EnvVars, error) {
	return env.ScanEnvVarsFromHost()
}

func (e *EnvVarManager) ScanDotEnvAndAttachEnvVars(dotEnvFile string,
	currentEnvVars types.EnvVars) (types.EnvVars, error) {
	if dotEnvFile == "" {
		return nil, errors.New("failed to get env vars from dot env file: file path is empty")
	}

	gitRepoDir, err := utils.FindGitRepoDir(3)
	if err != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file "+
			"cant find the root directory for the .git repository: %s", err.Error())
	}

	dotEnvFilePath := filepath.Join(gitRepoDir, dotEnvFile)

	if fileErr := utils.FileExistAndItIsAFile(dotEnvFilePath); fileErr != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file on path %s: %s",
			dotEnvFilePath, fileErr.Error())
	}

	if dotFileErr := utils.FileIsNotEmpty(dotEnvFilePath); dotFileErr != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file on path %s: %s, file is empty",
			dotEnvFilePath, dotFileErr.Error())
	}

	dotEnvVars, err := env.GetEnvVarsFromDotFile(env.DotEnvFileReadOpt{
		FilePath:       dotEnvFilePath,
		IgnoreComments: true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get env vars from dot env file on path %s: %s",
			dotEnvFilePath, err.Error())
	}

	return e.AttachEnvVarsToConfig(currentEnvVars, dotEnvVars)
}

func (e *EnvVarManager) AttachEnvVarsToConfig(envVarsCfg, envVars types.EnvVars) (types.EnvVars,
	error) {
	if err := env.SetEnvVars(envVars); err != nil {
		return nil, fmt.Errorf("failed to attach env vars to config: %s", err.Error())
	}

	return env.MergeEnvVars(envVarsCfg, envVars), nil
}

func NewEnvVarManager() EnvVarFetcher {
	return &EnvVarManager{}
}
