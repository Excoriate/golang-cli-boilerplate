package cli

import "github.com/Excoriate/golang-cli-boilerplate/pkg/types"

// IniDefaultFlagOptions represents the default flag options.
// This is used to initialize the CLI client.
type IniDefaultFlagOptions struct {
	DotEnvFile                   string
	ShowEnvVars                  bool
	ScanEnvVarsFromHost          bool
	ScanAWSEnvVarsFromHost       bool
	ScanTerraformEnvVarsFromHost bool
}

type AppCfg struct {
	// Env vars.
	EnvVarsHost      types.EnvVars
	EnvVarsAWS       types.EnvVars
	EnvVarsTerraform types.EnvVars
	EnvVarsDotEnd    types.EnvVars
	// FS related configuration.
	CurrentDir string
	HomeDir    string
	IsGitDir   bool
}
