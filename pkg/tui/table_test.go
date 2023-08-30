package tui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowTable(t *testing.T) {
	data := [][]string{
		{"1", "John", "25"},
		{"2", "Adam", "30"},
	}

	t.Run("Success", func(t *testing.T) {
		headers := []string{"#", "Name", "Age"}
		err := ShowTable(TableOptions{
			Headers: headers,
			Data:    data,
		})
		assert.NoError(t, err)
	})

	t.Run("Failure by Missmatch", func(t *testing.T) {
		headersMismatch := []string{"#", "Name"}
		err := ShowTable(TableOptions{
			Headers: headersMismatch,
			Data:    data,
		})
		assert.Error(t, err)

	})

	t.Run("Failure by Empty Headers", func(t *testing.T) {
		var emptyHeaders []string
		err := ShowTable(TableOptions{
			Headers: emptyHeaders,
			Data:    data,
		})
		assert.Error(t, err)
	})

	t.Run("Failure by Empty Data", func(t *testing.T) {
		var emptyData [][]string
		err := ShowTable(TableOptions{
			Headers: []string{"#", "Name", "Age"},
			Data:    emptyData,
		})
		assert.Error(t, err)
	})
}
