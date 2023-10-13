package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDoubleQuotes(t *testing.T) {
	t.Run("remove double quotes", func(t *testing.T) {
		str := RemoveDoubleQuotes(`"Hello, World!"`)
		assert.Equal(t, `Hello, World!`, str)
	})
	// Additional tests
}

func TestNormaliseStringToLower(t *testing.T) {
	t.Run("normalise to lower case", func(t *testing.T) {
		str := NormaliseStringToLower(`Hello, World!`)
		assert.Equal(t, `hello, world!`, str)
	})
	// Additional tests
}

func TestMergeSlices(t *testing.T) {
	t.Run("merge slices", func(t *testing.T) {
		res := MergeSlices([]string{"Hello"}, []string{"World"})
		assert.Equal(t, []string{"Hello", "World"}, res)
	})
	// Additional tests
}

// Repeat the process for the rest of the functions
