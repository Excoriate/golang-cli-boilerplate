package cli

import (
	"testing"

	"github.com/Excoriate/golang-cli-boilerplate/internal/config"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/types"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/o11y"
	"github.com/stretchr/testify/assert"
)

func TestWithScannedEnvVarsFromHost(t *testing.T) {
	b := Builder{
		initOptions: InitOptions{
			ScanEnvVarsFromHost: true,
		},
		logger: o11y.NewLogger(o11y.LoggerOptions{
			WriteToStdout: true,
		}),
		cliApp: &types.App{
			EnvVars: types.EnvVars{},
		},
		cfg: &config.Cfg{},
	}

	t.Run("Success", func(t *testing.T) {
		err := b.WithScannedEnvVarsFromHost()(&b.initOptions)

		assert.NoError(t, err, "The WithScannedEnvVarsFromHost should not return an error")
		assert.True(t, b.initOptions.ScanEnvVarsFromHost, "The WithScannedEnvVarsFromHost should set the ScanEnvVarsFromHost to true")

		// envVarsScanned
		scannedEnvVars := b.cliApp.EnvVars
		assert.NotNil(t, scannedEnvVars, "The WithScannedEnvVarsFromHost should set the EnvVars to a non-nil value")
		assert.NotEmpty(t, scannedEnvVars, "The WithScannedEnvVarsFromHost should set the EnvVars to a non-empty value")
	})
}
