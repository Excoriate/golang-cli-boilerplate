package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/cmd/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile             string
	debug               bool
	showEnvVars         bool
	scanEnvVarsFromHost bool
	dotEnvFile          string
	output              string
	save                bool
)

const CLIName = "golang-cli-boilerplate"
const CLIEnvVarPrefix = "AWS_ECS_DEPLOYER"

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     CLIName,
	Long: `ecs-deployer is a cmd-line tool that helps you manage your ECS services, and,
	related AWS infrastructure easily. It can be used in a stand-alone mode, or, as a
	library in your Go projects.`,
	Example: `
	  ecs-deployer deploy --service=myservice --cluster=mycluster`,
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
	rootCmd.PersistentFlags().BoolVarP(&debug,
		"debug",
		"d", false,
		"Enabled debug mode")

	rootCmd.PersistentFlags().StringVarP(&cfgFile,
		"config",
		"c", "",
		fmt.Sprintf("config file (default is $HOME/.%s.yaml)", CLIName))

	rootCmd.PersistentFlags().BoolVarP(&showEnvVars,
		"show-env-vars",
		"", false,
		"Show environment variables")

	rootCmd.PersistentFlags().StringVarP(&dotEnvFile,
		"dot-env-file",
		"", "",
		"DotEnv file that'll be used to load environment variables")

	rootCmd.PersistentFlags().BoolVarP(&scanEnvVarsFromHost,
		"scan-env-vars-from-host",
		"", false,
		"Scan environment variables from host")

	rootCmd.PersistentFlags().StringVarP(&output,
		"output",
		"o", "table",
		"Output format. One of: json|yaml/yml|table. If not provided, it defaults to table.")

	rootCmd.PersistentFlags().BoolVarP(&save,
		"save",
		"s", false,
		"Save output to file. If not provided, it defaults to stdout.")

	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("showEnvVars", rootCmd.PersistentFlags().Lookup("show-env-vars"))
	_ = viper.BindPFlag("dotEnvFile", rootCmd.PersistentFlags().Lookup("dot-env-file"))
	_ = viper.BindPFlag("scanEnvVarsFromHost", rootCmd.PersistentFlags().Lookup("scan-env-vars-from-host"))
	_ = viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	_ = viper.BindPFlag("save", rootCmd.PersistentFlags().Lookup("save"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(fmt.Sprintf(".%s", CLIName))

		_ = viper.SafeWriteConfig()
	}

	viper.SetEnvPrefix(CLIEnvVarPrefix)
	_ = viper.BindEnv("debug", "DEBUG")
	_ = viper.BindEnv("config", "CONFIG")
	_ = viper.BindEnv("showEnvVars", "SHOW_ENV_VARS")
	_ = viper.BindEnv("dotEnvFile", "DOT_ENV_FILE")
	_ = viper.BindEnv("scanEnvVarsFromHost", "SCAN_ENV_VARS_FROM_HOST")
	_ = viper.BindEnv("output", "OUTPUT")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.Set("CLI_NAME_TITLE", CLIName)
}

func addSubCommands() {
	rootCmd.AddCommand(aws.Cmd)
}

func init() {
	addPersistentFlagsToRootCMD()

	cobra.OnInitialize(initConfig)
	addSubCommands()
}
