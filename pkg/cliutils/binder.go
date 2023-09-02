package cliutils

import (
	"fmt"
	"strings"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/constants"
	"github.com/spf13/viper"
)

func BindAWSCreds() error {
	if err := viper.BindEnv("awsRegion", constants.AWSRegionEnvVar); err != nil {
		return err
	}

	if err := viper.BindEnv("awsAccessKeyID", constants.AWSAccessKeyEnvVar); err != nil {
		return err
	}

	return viper.BindEnv("awsSecretAccessKey", constants.AWSSecreetAccessKeyEnvVar)
}

func GetViperEnvVarPrefix(prefix string) string {
	prefixRaw := prefix
	// if prefix has -, replace them by _
	if strings.Contains(prefix, "-") {
		prefixRaw = strings.ReplaceAll(prefix, "-", "_")
	}

	return strings.TrimSpace(fmt.Sprintf("%s_", prefixRaw))
}
