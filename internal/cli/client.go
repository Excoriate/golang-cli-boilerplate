package cli

import (
	"context"
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/env"

	"github.com/Excoriate/golang-cli-boilerplate/internal/config"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"
)

type Builder struct {
	appCfg        *AppCfg
	envVarManager config.EnvVarFetcher
	logger        o11y.LoggerInterface
	ctx           context.Context
	ux            *UX
	initOptions   IniDefaultFlagOptions
}

type Client struct {
	App    *AppCfg
	Logger o11y.LoggerInterface
	Ctx    context.Context
	UX     *UX
}

type UX struct {
	Titles   tui.TitleWriter
	Messages tui.MessageWriter
}

type InitOptionsFunc func(*IniDefaultFlagOptions) error

func (b *Builder) Build(optFns ...func(*IniDefaultFlagOptions) error) (*Client, error) {
	for _, option := range optFns {
		if err := option(&b.initOptions); err != nil {
			return nil, err
		}
	}

	var client Client

	client.App = b.appCfg
	client.UX = b.ux
	client.Logger = b.logger
	client.Ctx = b.ctx

	return &client, nil
}

func (b *Builder) WithPrintedHostEnvVars() InitOptionsFunc {
	return func(o *IniDefaultFlagOptions) error {
		if !o.ShowEnvVars {
			return nil
		}

		envVars, err := b.envVarManager.ScanEnvVarsFromHost()
		if err != nil {
			return err
		}

		b.logger.Info(fmt.Sprintf("Scanned environment variables from host: %v", envVars))
		for k, v := range envVars {
			b.logger.Info(fmt.Sprintf("Scanned environment variable from host: %s=%s", k, v))
		}

		return nil
	}
}

func (b *Builder) WithScannedEnvVarsFromHost() InitOptionsFunc {
	return func(o *IniDefaultFlagOptions) error {
		if !o.ScanEnvVarsFromHost {
			return nil
		}

		envVars, err := b.envVarManager.ScanEnvVarsFromHost()
		if err != nil {
			return err
		}

		b.appCfg.EnvVarsHost = env.MergeEnvVars(b.appCfg.EnvVarsHost, envVars)
		return nil
	}
}

func (b *Builder) WithScannedAWSEnvVarsFromHost() InitOptionsFunc {
	return func(o *IniDefaultFlagOptions) error {
		if !o.ScanAWSEnvVarsFromHost {
			return nil
		}

		envVars, err := env.ScanAWSEnvVarsFromHost()
		if err != nil {
			return err
		}

		b.appCfg.EnvVarsAWS = envVars
		b.appCfg.EnvVarsHost = env.MergeEnvVars(b.appCfg.EnvVarsHost, envVars)
		return nil
	}
}

func (b *Builder) WithScannedTerraformEnvVarsFromHost() InitOptionsFunc {
	return func(o *IniDefaultFlagOptions) error {
		if !o.ScanTerraformEnvVarsFromHost {
			return nil
		}

		envVars, err := env.ScanTerraformEnvVarsFromHost()
		if err != nil {
			return err
		}

		b.appCfg.EnvVarsTerraform = envVars
		b.appCfg.EnvVarsHost = env.MergeEnvVars(b.appCfg.EnvVarsHost, envVars)
		return nil
	}
}

func (b *Builder) WithDotEnvFile() InitOptionsFunc {
	return func(o *IniDefaultFlagOptions) error {
		if o.DotEnvFile == "" {
			return nil
		}

		envVars, err := b.envVarManager.ScanDotEnvAndAttachEnvVars(o.DotEnvFile, b.appCfg.EnvVarsHost)
		if err != nil {
			return err
		}

		b.appCfg.EnvVarsDotEnd = envVars
		b.appCfg.EnvVarsHost = env.MergeEnvVars(b.appCfg.EnvVarsHost, envVars)
		return nil
	}
}

func (b *Builder) WithLogger() InitOptionsFunc {
	return func(o *IniDefaultFlagOptions) error {
		b.logger = o11y.NewLogger(o11y.LoggerOptions{
			EnableJSONHandler: true,
			EnableStdError:    true,
		})

		return nil
	}
}

func NewClient(ctx context.Context, options IniDefaultFlagOptions) *Builder {
	return &Builder{
		appCfg: &AppCfg{
			EnvVarsHost:      make(map[string]string), // Initialize the map.
			EnvVarsTerraform: map[string]string{},
			EnvVarsAWS:       map[string]string{},
			HomeDir:          utils.GetHomeDir(),
			CurrentDir:       utils.GetCurrentDir(),
		},
		envVarManager: config.NewEnvVarManager(),
		logger:        nil,
		initOptions:   options,
		ctx:           ctx,
		ux: &UX{
			Titles:   tui.NewTitleWriter(),
			Messages: tui.NewMessageWriter(),
		},
	}
}
