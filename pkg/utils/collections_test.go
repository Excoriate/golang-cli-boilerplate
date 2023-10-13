package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStrInSlice(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		slice := []string{"value1", "value2", "value3"}
		str, err := FindStrInSlice(slice, "value1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", str)
	})

	t.Run("not found", func(t *testing.T) {
		slice := []string{"value1", "value2", "value3"}
		_, err := FindStrInSlice(slice, "value4")
		assert.Error(t, err)
	})
}

func TestFindStrInMapValues(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
		value, err := FindStrInMapValues(m, "value1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", value)
	})

	t.Run("not found", func(t *testing.T) {
		m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
		_, err := FindStrInMapValues(m, "value4")
		assert.Error(t, err)
	})
}

func TestFindStrInMapKeys(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
		key, err := FindStrInMapKeys(m, "key1")
		assert.NoError(t, err)
		assert.Equal(t, "key1", key)
	})

	t.Run("not found", func(t *testing.T) {
		m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
		_, err := FindStrInMapKeys(m, "key4")
		assert.Error(t, err)
	})
}

func TestKeyInMapHasValueSet(t *testing.T) {
	t.Run("value exists", func(t *testing.T) {
		m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
		err := KeyInMapHasValueSet(m, "key1")
		assert.NoError(t, err)
	})

	t.Run("value is empty", func(t *testing.T) {
		m := map[string]string{"key1": "", "key2": "value2", "key3": "value3"}
		err := KeyInMapHasValueSet(m, "key1")
		assert.Error(t, err)
	})
}
