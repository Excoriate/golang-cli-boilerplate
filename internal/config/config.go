package config

import (
	"fmt"
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/env"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/errs"
	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
)

type EnvironmentScanner interface {
	ScanEnvVarsFromHost() (types.EnvVars, error)
	GetEnvVarsFromDotEnvFile(dotEnvFile string, currentEnvVars types.EnvVars) (types.EnvVars, error)
}

type Cfg struct{}

// GetEnvVarsFromDotEnvFile returns the env vars from the dot env file
// If it success, it'll attach the found environment variables to the current configuration
// If it fails, it'll return an error
func (c *Cfg) GetEnvVarsFromDotEnvFile(dotEnvFile string, currentEnvVars types.EnvVars) (types.
	EnvVars, error) {
	envVarsFromDotEnv, err := GetEnvVarsFromDotEnv(dotEnvFile)

	if err != nil {
		return nil, errs.NewConfigurationErr(errs.Opts{
			Error:   err,
			Details: fmt.Sprintf("Failed to get env vars from dot env file: %s", dotEnvFile),
		})
	}

	// Attach to existing env vars
	envVars, err := AttachEnvVarsToConfig(currentEnvVars, envVarsFromDotEnv)
	if err != nil {
		return nil, errs.NewConfigurationErr(errs.Opts{
			Error:   err,
			Details: "failed to attach env vars to config",
		})
	}

	return envVars, nil
}

// ScanEnvVarsFromHost returns the env vars from the host
// If it success, it'll attach the found environment variables to the current configuration
// If it fails, it'll return an error
func (c *Cfg) ScanEnvVarsFromHost() (types.EnvVars, error) {
	envVarsScanned, err := env.GetAllFromHost()
	if err != nil {
		return nil, errs.NewConfigurationErr(errs.Opts{
			Error:   err,
			Details: "Failed to get env vars from host",
		})
	}

	return envVarsScanned, nil
}

func NewCLIApp(logger o11y.LoggerInterface) *types.App {
	cfg := &types.App{
		EnvVars: types.EnvVars{},
	}

	homeDir, _ := os.UserHomeDir()
	currentDir, _ := os.Getwd()

	logger.Info("Home directory", o11y.LoggerArg{
		Key:   "homeDir",
		Value: homeDir,
	})
	logger.Info("Current directory", o11y.LoggerArg{
		Key:   "currentDir",
		Value: currentDir,
	})

	cfg.CurrentDir = currentDir
	cfg.HomeDir = homeDir

	_, err := utils.FindGitRepoDir(3)
	if err != nil {
		logger.Warn("this is not a .git repo")
		cfg.IsGitDir = false
	}

	cfg.IsGitDir = true

	return cfg
}
