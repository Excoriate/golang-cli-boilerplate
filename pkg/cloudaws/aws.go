package cloudaws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
)

// NewAWSDefaultConfig returns a new aws.Config with the default credentials
// It is a wrapper around awsCfg.LoadDefaultConfig
func NewAWSDefaultConfig(ctx context.Context, region string) (aws.Config, error) {
	awsAuth, err := awsCfg.LoadDefaultConfig(ctx, awsCfg.WithRegion(region))
	if err != nil {
		return aws.Config{}, err
	}

	return awsAuth, nil
}

var awsCredKeys = [6]string{
	"aws_access_key_id",
	"aws_secret_access_key",
	"aws_session_token",
	"aws_access_key",
	"aws_secret_key",
	"aws_security_token",
}

type AWSCred struct {
	Key   string
	Value string
	IsSet bool
}

func DetectAWSCredentialsExported() bool {
	awsEnvVars := GetAllExportedAWSCredentials()

	return len(awsEnvVars) > 0
}

func GetAllExportedAWSCredentials() map[string]string {
	awsCreds := make(map[string]string)

	for _, key := range awsCredKeys {
		if val, exists := os.LookupEnv(key); exists {
			awsCreds[key] = val
		}
	}
	return awsCreds
}

func GetAWSCredentialsReport() []AWSCred {
	awsCreds := GetAllExportedAWSCredentials()

	var awsCredsReport []AWSCred
	for key, val := range awsCreds {
		awsCredsReport = append(awsCredsReport, AWSCred{
			Key:   key,
			Value: val,
			IsSet: true,
		})
	}

	return awsCredsReport
}

func GetAWSCredentialsValueByKey(key string) (*AWSCred, error) {
	if key == "" {
		return nil, fmt.Errorf("key is empty")
	}

	awsCreds := GetAllExportedAWSCredentials()

	if len(awsCreds) == 0 {
		return nil, fmt.Errorf("no AWS credentials found in environment variables")
	}

	if val, ok := awsCreds[key]; ok {
		return &AWSCred{
			Key:   key,
			Value: val,
			IsSet: true,
		}, nil
	}
	return nil, fmt.Errorf("AWS credential with key '%s' not found in environment variables", key)
}

func IsAWSCredSet(key string) error {
	if key == "" {
		return fmt.Errorf("key is empty")
	}

	awsCredsReport := GetAWSCredentialsReport()

	for _, cred := range awsCredsReport {
		if cred.Key == key && cred.IsSet {
			return nil
		}

		if cred.Key == key && !cred.IsSet {
			return fmt.Errorf("AWS credential with key '%s' is not set", key)
		}
	}

	return fmt.Errorf("AWS credential with key '%s' not found in environment variables", key)
}
