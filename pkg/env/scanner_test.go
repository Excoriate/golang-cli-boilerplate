package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanEnvVarsFromHost(t *testing.T) {
	t.Run("scan env vars", func(t *testing.T) {
		_ = os.Setenv("TEST_MY_ENV_VAR1", `"env_value1"`)
		_ = os.Setenv("TEST_MY_ENV_VAR2", `"env_value2"`)
		_ = os.Setenv("TEST_MY_ENV_VAR3", `"env_value3"`)

		res, err := ScanEnvVarsFromHost()
		require.NoError(t, err)
		require.Equal(t, "env_value1", res["TEST_MY_ENV_VAR1"])
		require.Equal(t, "env_value2", res["TEST_MY_ENV_VAR2"])
		require.Equal(t, "env_value3", res["TEST_MY_ENV_VAR3"])

		_ = os.Unsetenv("TEST_MY_ENV_VAR1")
		_ = os.Unsetenv("TEST_MY_ENV_VAR2")
		_ = os.Unsetenv("TEST_MY_ENV_VAR3")
	})
}

func TestScanEnvVarsWithPrefix(t *testing.T) {
	t.Run("scan env vars with prefix", func(t *testing.T) {
		_ = os.Setenv("PREFIX_MY_ENV_VAR1", `"prefix_env_value1"`)
		_ = os.Setenv("PREFIX_MY_ENV_VAR2", `"prefix_env_value2"`)
		_ = os.Setenv("TEST_MY_ENV_VAR3", `"env_value3"`)

		res, err := ScanEnvVarsWithPrefix("PREFIX")
		require.NoError(t, err)
		require.Equal(t, "prefix_env_value1", res["PREFIX_MY_ENV_VAR1"])
		require.Equal(t, "prefix_env_value2", res["PREFIX_MY_ENV_VAR2"])
		require.Emptyf(t, res["TEST_MY_ENV_VAR3"], "should not return env var with prefix")

		_ = os.Unsetenv("PREFIX_MY_ENV_VAR1")
		_ = os.Unsetenv("PREFIX_MY_ENV_VAR2")
		_ = os.Unsetenv("TEST_MY_ENV_VAR3")
	})
}
