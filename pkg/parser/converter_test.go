package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlToStructFromFile(t *testing.T) {
	t.Run("should return an error if the file does not exist", func(t *testing.T) {
		err := YamlToStructFromFile("non-existent-file", nil)

		assert.Errorf(t, err, "could not open the yaml file: open non-existent-file: no such file or directory")
	})
}
