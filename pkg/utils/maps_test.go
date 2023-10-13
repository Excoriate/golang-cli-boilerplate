package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapIsNulOrEmpty(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		assert.Equal(t, true, MapIsNulOrEmpty(map[string]string{}))
	})

	t.Run("non-empty map", func(t *testing.T) {
		assert.Equal(t, false, MapIsNulOrEmpty(map[string]string{"key": "value"}))
	})
}

func TestValueInMapIsNotEmptyByKey(t *testing.T) {
	t.Run("key exists in map", func(t *testing.T) {
		value, err := ValueInMapIsNotEmptyByKey(map[string]string{"key": "value"}, "key")
		assert.Equal(t, "value", value)
		assert.NoError(t, err)
	})

	t.Run("key does not exist in map", func(t *testing.T) {
		_, err := ValueInMapIsNotEmptyByKey(map[string]string{"key": "value"}, "nokey")
		assert.Error(t, err)
	})
	// Other cases
}
