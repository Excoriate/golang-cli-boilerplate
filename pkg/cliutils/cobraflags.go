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
	DefaultValue             interface{}
	FlagType                 string
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
		if option.FlagType == "" {
			option.FlagType = "string"
		}

		switch option.FlagType {
		case "string":
			if option.DefaultValue == nil {
				option.DefaultValue = ""
			}
			if option.IsPersistent {
				cmd.PersistentFlags().StringP(option.LongName, option.ShortName, option.DefaultValue.(string), option.Usage)
				flag = cmd.PersistentFlags().Lookup(option.LongName)
			} else {
				cmd.Flags().StringP(option.LongName, option.ShortName, option.DefaultValue.(string), option.Usage)
				flag = cmd.Flags().Lookup(option.LongName)
			}
		case "bool":
			if option.DefaultValue == nil {
				option.DefaultValue = false
			}
			if option.IsPersistent {
				cmd.PersistentFlags().BoolP(option.LongName, option.ShortName, option.DefaultValue.(bool), option.Usage)
				flag = cmd.PersistentFlags().Lookup(option.LongName)
			} else {
				cmd.Flags().BoolP(option.LongName, option.ShortName, option.DefaultValue.(bool), option.Usage)
				flag = cmd.Flags().Lookup(option.LongName)
			}
		case "int":
			if option.DefaultValue == nil {
				option.DefaultValue = 0
			}
			if option.IsPersistent {
				cmd.PersistentFlags().IntP(option.LongName, option.ShortName, option.DefaultValue.(int), option.Usage)
				flag = cmd.PersistentFlags().Lookup(option.LongName)
			} else {
				cmd.Flags().IntP(option.LongName, option.ShortName, option.DefaultValue.(int), option.Usage)
				flag = cmd.Flags().Lookup(option.LongName)
			}
		case "map":
			if option.DefaultValue == nil {
				option.DefaultValue = make(map[string]string)
			}
			if option.IsPersistent {
				cmd.PersistentFlags().StringToStringP(option.LongName, option.ShortName, option.DefaultValue.(map[string]string), option.Usage)
				flag = cmd.PersistentFlags().Lookup(option.LongName)
			} else {
				cmd.Flags().StringToStringP(option.LongName, option.ShortName, option.DefaultValue.(map[string]string), option.Usage)
				flag = cmd.Flags().Lookup(option.LongName)
			}
		case "slice":
			if option.DefaultValue == nil {
				option.DefaultValue = make([]string, 0)
			}

			if option.IsPersistent {
				cmd.PersistentFlags().StringSliceP(option.LongName, option.ShortName, option.DefaultValue.([]string), option.Usage)
				flag = cmd.PersistentFlags().Lookup(option.LongName)
			} else {
				cmd.Flags().StringSliceP(option.LongName, option.ShortName, option.DefaultValue.([]string), option.Usage)
				flag = cmd.Flags().Lookup(option.LongName)
			}
		default:
			return fmt.Errorf("unsupported flag type %s", option.FlagType)
		}

		if option.ViperBindingCfg.IsEnabled {
			if option.ViperBindingCfg.EnvVariableNameInViper == "" {
				return fmt.Errorf("env variable name in viper is required when the viper binding is enabled")
			}

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
