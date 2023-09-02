package cloudaws

import (
	"context"
	"os"
	"strings"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/errs"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type InitAWSAdapterOptions struct {
	Region                string
	SharedCredentialsFile string
	Profile               string
	Creds                 AWSCreds
}

type AWSCreds struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

type InitAWSAdapterOptionsFunc func(*InitAWSAdapterOptions) error

type AWSAdapter struct {
	Region string
	// Add clients here.
	ECSClient            *ecs.Client
	ECRClient            *ecr.Client
	CloudWatchLogsClient *cloudwatchlogs.Client
	Logger               o11y.LoggerInterface
}

type Builder struct {
	region                string
	sharedCredentialsFile string
	profile               string
	logger                o11y.LoggerInterface
	// The one that's injected into each specific client.
	adapter aws.Config
	// Clients.
	ecsClient            *ecs.Client
	ecrClient            *ecr.Client
	cloudWatchLogsClient *cloudwatchlogs.Client
}

func (b *Builder) Build(optFns ...func(*InitAWSAdapterOptions) error) (*AWSAdapter, error) {
	for _, option := range optFns {
		if err := option(&InitAWSAdapterOptions{
			Region:                b.region,
			SharedCredentialsFile: b.sharedCredentialsFile,
			Profile:               b.profile,
		}); err != nil {
			return nil, err
		}
	}

	var adapter AWSAdapter
	adapter.Region = b.region
	adapter.Logger = b.logger

	// Clients built
	adapter.ECSClient = b.ecsClient
	adapter.ECRClient = b.ecrClient
	adapter.CloudWatchLogsClient = b.cloudWatchLogsClient

	return &adapter, nil
}

type InitECSClientOptions struct {
}

func (b *Builder) WithECS() func(*InitAWSAdapterOptions) error {
	return func(options *InitAWSAdapterOptions) error {
		client := ecs.NewFromConfig(b.adapter)
		b.ecsClient = client

		return nil
	}
}

func (b *Builder) WithECR() func(*InitAWSAdapterOptions) error {
	return func(options *InitAWSAdapterOptions) error {
		client := ecr.NewFromConfig(b.adapter)
		b.ecrClient = client

		return nil
	}
}

func (b *Builder) WithCloudWatchLogs() func(*InitAWSAdapterOptions) error {
	return func(options *InitAWSAdapterOptions) error {
		client := cloudwatchlogs.NewFromConfig(b.adapter)
		b.cloudWatchLogsClient = client
		return nil
	}
}

func NewClient(ctx context.Context, l o11y.LoggerInterface, options InitAWSAdapterOptions) (*Builder,
	error) {
	logger := l
	if options.Region == "" {
		return nil, errs.NewAdapterErr(errs.Opts{
			Details: "failed to create a new AWS adapter instance, region is empty",
		})
	}

	if options.Creds.AccessKeyID != "" {
		logger.Info("The 'AWS access key ID was set explicitly. " +
			"It'll take precedence over the one set in environment variables.")
		_ = os.Setenv("AWS_ACCESS_KEY_ID", strings.TrimSpace(options.Creds.AccessKeyID))
	} else {
		if err := IsAWSCredSet("aws_access_key_id"); err != nil {
			return nil, errs.NewAdapterErr(errs.Opts{
				Details: "failed to create a new AWS adapter instance, error occurred while checking if AWS_ACCESS_KEY_ID is set",
				Error:   err,
			})
		}
	}

	if options.Creds.SecretAccessKey != "" {
		logger.Info("The 'AWS secret access key' was set explicitly. " +
			"It'll take precedence over the one set in environment variables.")
		_ = os.Setenv("AWS_SECRET_ACCESS_KEY", strings.TrimSpace(options.Creds.SecretAccessKey))
	} else {
		if err := IsAWSCredSet("aws_secret_access_key"); err != nil {
			return nil, errs.NewAdapterErr(errs.Opts{
				Details: "failed to create a new AWS adapter instance, error occurred while checking if AWS_SECRET_ACCESS_KEY is set",
				Error:   err,
			})
		}
	}

	adapter, err := NewAWSDefaultConfig(ctx, options.Region)
	if err != nil {
		return nil, errs.NewAdapterErr(errs.Opts{
			Details: "failed to create a new AWS adapter instance, error occurred while creating AWS default config",
			Error:   err,
		})
	}

	return &Builder{
		// Options required to resolve which configuration to use.
		region:                options.Region,
		sharedCredentialsFile: options.SharedCredentialsFile,
		profile:               options.Profile,
		// General purpose logger.
		logger: l,
		// Adapter, or connector.
		adapter: adapter,
	}, nil
}
