package env

import "github.com/spf13/viper"

func BindAWSKeysToViper() error {
	if err := viper.BindEnv("awsAccessKeyID", "AWS_ACCESS_KEY_ID"); err != nil {
		return err
	}

	return viper.BindEnv("awsSecretAccessKey", "AWS_SECRET_ACCESS_KEY")
}
