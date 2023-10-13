package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractBearerValueFromHeader(t *testing.T) {
	t.Run("valid header", func(t *testing.T) {
		token, err := ExtractBearerValueFromHeader("Bearer token123")
		require.NoError(t, err)
		assert.Equal(t, "token123", token)
	})

	t.Run("missing token value", func(t *testing.T) {
		_, err := ExtractBearerValueFromHeader("Bearer ")
		require.Error(t, err)
	})

	t.Run("missing bearer type", func(t *testing.T) {
		_, err := ExtractBearerValueFromHeader("token123")
		require.Error(t, err)
	})
}

func TestValidateAuthorizationHeader(t *testing.T) {
	t.Run("valid header", func(t *testing.T) {
		headers := map[string]string{"Authorization": "Bearer token123", "Other": "value"}
		err := ValidateAuthorizationHeader(headers)
		require.NoError(t, err)
	})

	t.Run("missing Authorization", func(t *testing.T) {
		headers := map[string]string{"Other": "value"}
		err := ValidateAuthorizationHeader(headers)
		require.Error(t, err)
	})

	t.Run("missing token value", func(t *testing.T) {
		headers := map[string]string{"Authorization": "", "Other": "value"}
		err := ValidateAuthorizationHeader(headers)
		require.Error(t, err)
	})
}

func TestValidateBearerTokenSchema(t *testing.T) {
	t.Run("valid schema", func(t *testing.T) {
		err := ValidateBearerTokenSchema("Bearer token123")
		require.NoError(t, err)
	})

	t.Run("missing token value", func(t *testing.T) {
		err := ValidateBearerTokenSchema("Bearer ")
		require.Error(t, err)
	})

	t.Run("missing bearer type", func(t *testing.T) {
		err := ValidateBearerTokenSchema("token123")
		require.Error(t, err)
	})
}
