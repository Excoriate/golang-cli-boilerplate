package cli

import (
	"context"
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/internal/config"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/tui"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"
)

type Builder struct {
	cliApp      *types.App
	cfg         config.EnvironmentScanner
	logger      o11y.LoggerInterface
	ctx         context.Context
	ux          *UX
	initOptions InitOptions
}

type Client struct {
	App    *types.App
	Logger o11y.LoggerInterface
	Ctx    context.Context
	UX     *UX
}

type UX struct {
	Titles   tui.UXTitler
	Messages tui.UXMessage
}

type InitOptions struct {
	DotEnvFile          string
	ShowEnvVars         bool
	ScanEnvVarsFromHost bool
}

type InitOptionsFunc func(*InitOptions) error

func (b *Builder) Build(optFns ...func(*InitOptions) error) (*Client, error) {
	for _, option := range optFns {
		if err := option(&b.initOptions); err != nil {
			return nil, err
		}
	}

	var client Client

	client.App = b.cliApp
	client.UX = b.ux
	client.Logger = b.logger
	client.Ctx = b.ctx

	return &client, nil
}

func (b *Builder) WithPrintedHostEnvVars() InitOptionsFunc {
	return func(o *InitOptions) error {
		if !o.ShowEnvVars {
			return nil
		}

		envVars, err := b.cfg.ScanEnvVarsFromHost()
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
	return func(o *InitOptions) error {
		if !o.ScanEnvVarsFromHost {
			return nil
		}

		envVars, err := b.cfg.ScanEnvVarsFromHost()
		if err != nil {
			return err
		}

		b.cliApp.EnvVars = envVars
		return nil
	}
}

func (b *Builder) WithDotEnvFile() InitOptionsFunc {
	return func(o *InitOptions) error {
		if o.DotEnvFile == "" {
			return nil
		}

		envVars, err := b.cfg.GetEnvVarsFromDotEnvFile(o.DotEnvFile, b.cliApp.EnvVars)
		if err != nil {
			return err
		}

		b.cliApp.EnvVars = envVars
		return nil
	}
}

func New(ctx context.Context, options InitOptions) *Builder {
	logger := o11y.NewLogger(o11y.LoggerOptions{
		RegisterCallerFunction: true,
		WriteToStdout:          false, // It'll be written to stderr.
	})

	// Load basic configuration.
	cliApp := config.NewCLIApp(logger)

	return &Builder{
		cliApp:      cliApp,
		cfg:         &config.Cfg{},
		logger:      logger,
		initOptions: options,
		ctx:         ctx,
		ux: &UX{
			Titles:   tui.NewTitler(),
			Messages: tui.NewMessage(),
		},
	}
}
