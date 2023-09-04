package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/internal/cli"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/cliutils"

	"github.com/Excoriate/golang-cli-boilerplate/cmd/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

const CLIName = "golang-cli-boilerplate"

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     CLIName,
	Long: fmt.Sprintf(`%s is a cmd-line tool that helps you manage your ECS services, and,
	related AWS infrastructure easily. It can be used in a stand-alone mode, or, as a
	library in your Go projects.`, CLIName),
	Example: fmt.Sprintf(`
	  %s deploy --service=myservice --cluster=mycluster`, CLIName),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize the CLI client.
		client, err := cli.New()
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.WithValue(context.Background(), "client", client)
		cmd.SetContext(ctx)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		os.Exit(1)
	}
}

func addPersistentFlagsToRootCMD() {
	_ = cliutils.AddFlags(rootCmd, []cliutils.CobraFlagOptions{
		{
			ShortName:                "d",
			LongName:                 "debug",
			Usage:                    "Enable debug mode",
			FlagType:                 "bool",
			DefaultValue:             false,
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "debug",
				EnvVariableNameInHost:  "DEBUG",
				IsEnabled:              true,
			},
		},
		{
			ShortName:                "c",
			LongName:                 "config",
			Usage:                    fmt.Sprintf("config file (default is $HOME/.%s.yaml)", CLIName),
			DefaultValue:             "",
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "config",
				EnvVariableNameInHost:  "CONFIG",
				IsEnabled:              true,
			},
		},
		{
			LongName:                 "show-env-vars",
			Usage:                    "Show environment variables",
			FlagType:                 "bool",
			DefaultValue:             false,
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "showEnvVars",
				EnvVariableNameInHost:  "SHOW_ENV_VARS",
				IsEnabled:              true,
			},
		},
		{
			LongName:                 "dot-env-file",
			Usage:                    "DotEnv file that'll be used to load environment variables",
			DefaultValue:             "",
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "dotEnvFile",
				EnvVariableNameInHost:  "DOT_ENV_FILE",
				IsEnabled:              true,
			},
		},
		{
			LongName:                 "scan-env-vars-from-host",
			Usage:                    "Scan environment variables from host",
			FlagType:                 "bool",
			DefaultValue:             false,
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "scanEnvVarsFromHost",
				EnvVariableNameInHost:  "SCAN_ENV_VARS_FROM_HOST",
				IsEnabled:              true,
			},
		},
		{
			LongName:                 "scan-terraform-env-vars",
			Usage:                    "Scan Terraform specific environment variables from host",
			FlagType:                 "bool",
			DefaultValue:             false,
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "scanTerraformEnvVarsFromHost",
				EnvVariableNameInHost:  "SCAN_TERRAFORM_ENV_VARS_FROM_HOST",
				IsEnabled:              true,
			},
		},
		{
			LongName:                 "scan-aws-env-vars",
			Usage:                    "Scan AWS specific environment variables from host",
			FlagType:                 "bool",
			DefaultValue:             false,
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "scanAWSEnvVarsFromHost",
				EnvVariableNameInHost:  "SCAN_AWS_ENV_VARS_FROM_HOST",
				IsEnabled:              true,
			},
		},
		{
			ShortName:                "o",
			LongName:                 "output",
			Usage:                    "Output format. One of: json|yaml/yml|table. If not provided, it defaults to table.",
			DefaultValue:             "table",
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "output",
				EnvVariableNameInHost:  "OUTPUT",
				IsEnabled:              true,
			},
		},
		{
			ShortName:                "s",
			LongName:                 "save",
			Usage:                    "Save output to file. If not provided, it defaults to stdout.",
			FlagType:                 "bool",
			DefaultValue:             false,
			IsPersistent:             true,
			FailIfEnvVarBindingFails: false,
			ViperBindingCfg: cliutils.CobraViperBindingOptions{
				EnvVariableNameInViper: "save",
				EnvVariableNameInHost:  "SAVE",
				IsEnabled:              true,
			},
		},
	})
}

func initConfig() {
	// Initialize the viper configuration.
	if err := cliutils.InitViperConfig(cliutils.ViperConfigOptions{
		CfgName:       CLIName,
		CfgFile:       cfgFile,
		EnvVarsPrefix: cliutils.GetViperEnvVarPrefix(CLIName),
	}); err != nil {
		log.Fatal(err)
	}

	// Expose the CLI name to the Viper configuration.
	viper.Set("CLI_NAME_TITLE", CLIName)
}

func init() {
	addPersistentFlagsToRootCMD()
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(aws.Cmd)
}
