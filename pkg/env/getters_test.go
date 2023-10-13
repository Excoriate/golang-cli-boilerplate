package env

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFromMultiplePotentialKeys(t *testing.T) {
	t.Run("no keys allowed", func(t *testing.T) {
		_, err := GetFromMultiplePotentialKeys(nil)
		require.Error(t, err)
	})

	t.Run("key exists and value set", func(t *testing.T) {
		_ = os.Setenv("MY_EXISTING_KEY", `"env_value1"`)
		res, err := GetFromMultiplePotentialKeys([]string{"MY_EXISTING_KEY"})
		require.NoError(t, err)
		assert.Equal(t, VarSkeleton{"MY_EXISTING_KEY", `"env_value1"`}, res)

		_ = os.Unsetenv("MY_EXISTING_KEY")
	})

	t.Run("multiple keys exist", func(t *testing.T) {
		keys := []string{"KEY1", "KEY2", "KEY3"}
		for _, key := range keys {
			_ = os.Setenv(key, fmt.Sprintf(`"env_value_%s"`, strings.TrimPrefix(key, "KEY")))
		}
		res, err := GetFromMultiplePotentialKeys(keys)
		require.NoError(t, err)
		assert.Equal(t, VarSkeleton{"KEY1", `"env_value_1"`}, res)

		for _, key := range keys {
			_ = os.Unsetenv(key)
		}
	})

	t.Run("key exists but value not set", func(t *testing.T) {
		_ = os.Setenv("MY_EMPTY_KEY", "")
		_, err := GetFromMultiplePotentialKeys([]string{"MY_EMPTY_KEY"})
		require.Error(t, err)

		_ = os.Unsetenv("MY_EMPTY_KEY")
	})

	t.Run("key does not exist", func(t *testing.T) {
		_, err := GetFromMultiplePotentialKeys([]string{"NON_EXISTING_KEY"})
		require.Error(t, err)
	})
}
