package cliutils

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CobraFlagOptions struct {
	ShortName                string
	LongName                 string
	Usage                    string
	DefaultValue             string
	FailIfEnvVarBindingFails bool
	IsPersistent             bool
	ViperBindingCfg          CobraViperBindingOptions
	IsRequired               bool
}

type CobraViperBindingOptions struct {
	EnvVariableNameInViper string
	EnvVariableNameInHost  string
	IsEnabled              bool
}

func AddFlags(cmd *cobra.Command, flags []CobraFlagOptions) error {
	for _, option := range flags {
		var flag *pflag.Flag
		if option.IsPersistent {
			cmd.PersistentFlags().StringP(option.LongName, option.ShortName, option.DefaultValue, option.Usage)
			flag = cmd.PersistentFlags().Lookup(option.LongName)
		} else {
			cmd.Flags().StringP(option.LongName, option.ShortName, option.DefaultValue, option.Usage)
			flag = cmd.Flags().Lookup(option.LongName)
		}

		if option.ViperBindingCfg.IsEnabled {
			if err := viper.BindPFlag(option.ViperBindingCfg.EnvVariableNameInViper, flag); err != nil {
				return fmt.Errorf("failed to bind flag %s to viper: %w", option.ViperBindingCfg.EnvVariableNameInViper, err)
			}
		}

		if option.ViperBindingCfg.EnvVariableNameInHost != "" {
			err := viper.BindEnv(option.ViperBindingCfg.EnvVariableNameInViper, option.ViperBindingCfg.EnvVariableNameInHost)
			if err != nil && option.FailIfEnvVarBindingFails {
				return fmt.Errorf("failed to bind env var %s to viper: %w", option.ViperBindingCfg.EnvVariableNameInHost, err)
			}
		}

		// Check if flag is required and is not set
		if option.IsRequired {
			if err := markAsRequiredAndCheck(cmd, option.LongName, option.IsPersistent); err != nil {
				return err
			}
		}
	}

	return nil
}

func markAsRequiredAndCheck(cmd *cobra.Command, flagName string, isPersistent bool) error {
	if isPersistent {
		if err := cmd.MarkPersistentFlagRequired(flagName); err != nil {
			return fmt.Errorf("failed to mark flag %s as required: %w", flagName, err)
		}
	} else {
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			return fmt.Errorf("failed to mark flag %s as required: %w", flagName, err)
		}
	}

	return nil
}