package env

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanAWSEnvVarsFromHost(t *testing.T) {
	t.Run("Scan existing AWS env vars", func(t *testing.T) {
		_ = os.Setenv("AWS_FAKE", "fake")

		envVars, err := ScanAWSEnvVarsFromHost()
		assert.NoErrorf(t, err, "ScanAWSEnvVarsFromHost should not return an error")
		assert.NotNil(t, envVars, "ScanAWSEnvVarsFromHost should return a map")
		assert.NotEmpty(t, envVars, "ScanAWSEnvVarsFromHost should return a map with values")
		assert.Equal(t, "fake", envVars["AWS_FAKE"], "ScanAWSEnvVarsFromHost should return a map with the correct values")
	})

	t.Run("There is no AWS env vars", func(t *testing.T) {
		// Unset all env vars that starts with AWS_
		for _, env := range os.Environ() {
			// If starts with AWS_, unset it.
			if strings.HasPrefix(env, "AWS_") {
				pair := strings.SplitN(env, "=", 2)
				key := pair[0]
				_ = os.Unsetenv(key)
			}
		}

		envVars, err := ScanAWSEnvVarsFromHost()
		assert.Errorf(t, err, "ScanAWSEnvVarsFromHost should return an error")
		assert.Nil(t, envVars, "ScanAWSEnvVarsFromHost should return nil")
	})
}
