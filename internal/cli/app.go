package cli

import (
	"context"

	"github.com/spf13/viper"
)

func New() (*Client, error) {
	client := NewClient(context.TODO(), IniDefaultFlagOptions{
		DotEnvFile:                   viper.GetString("dotEnvFile"),
		ScanEnvVarsFromHost:          viper.GetBool("scanEnvVarsFromHost"),
		ShowEnvVars:                  viper.GetBool("showEnvVars"),
		ScanAWSEnvVarsFromHost:       viper.GetBool("scanAWSEnvVarsFromHost"),
		ScanTerraformEnvVarsFromHost: viper.GetBool("scanTerraformEnvVarsFromHost"),
	})

	// Build the CLI client.
	cliClient, err := client.Build(
		client.WithLogger(),
		client.WithScannedEnvVarsFromHost(),
		client.WithDotEnvFile(),
		client.WithScannedTerraformEnvVarsFromHost(),
		client.WithScannedAWSEnvVarsFromHost(),
		client.WithPrintedHostEnvVars())

	if err != nil {
		return nil, err
	}

	return cliClient, nil
}
