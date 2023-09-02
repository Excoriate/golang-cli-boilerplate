package cliutils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ViperConfigOptions struct {
	CfgFile       string
	CfgName       string
	CfgFileType   string
	EnvVarsPrefix string
}

func InitViperConfig(options ViperConfigOptions) error {
	if options.CfgFileType == "" {
		options.CfgFileType = "yaml"
	}

	if options.CfgName == "" {
		return fmt.Errorf("failed to configure viper: config name is empty")
	}

	if options.CfgFile != "" {
		viper.SetConfigFile(options.CfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get user home directory: %w", err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType(options.CfgFileType)
		viper.SetConfigName(fmt.Sprintf(".%s", options.CfgName))

		// Consider deleting or revising SafeWriteConfig here.

		_ = viper.SafeWriteConfig()
	}

	if options.EnvVarsPrefix != "" {
		viper.SetEnvPrefix(options.EnvVarsPrefix)
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the config file: %w", err)
	}

	return nil
}
