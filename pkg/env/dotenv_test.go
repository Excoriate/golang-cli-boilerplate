package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDotEnvLoad(t *testing.T) {
	t.Run("error for non-existing file name", func(t *testing.T) {
		dotEnv := NewDotEnv()

		mockUtils := &MockedUtils{}
		mockUtils.FileExistAndItIsAFile = mockUtils.FileExistAndIsAFileMock

		err := dotEnv.LoadByName("non_existing_env")

		assert.Error(t, err)
	})

	t.Run("load existing file by its name", func(t *testing.T) {
		dotEnv := NewDotEnv()
		fakeNewEnvFile()

		mockUtils := &MockedUtils{}
		mockUtils.FileExistAndItIsAFile = mockUtils.FileExistAndIsAFileMock

		err := dotEnv.LoadByName(".env.example")

		assert.NoError(t, err)
		cleanFakeFile()
	})

	t.Run("load non-existing env file", func(t *testing.T) {
		dotEnv := NewDotEnv()

		mockUtils := &MockedUtils{}
		mockUtils.FileExistAndItIsAFile = mockUtils.FileExistAndIsAFileMock

		err := dotEnv.LoadByEnv("non_existing_env")

		assert.Error(t, err)
	})

	t.Run("load env file without env provided", func(t *testing.T) {
		dotEnv := NewDotEnv()

		mockUtils := &MockedUtils{}
		mockUtils.FileExistAndItIsAFile = mockUtils.FileExistAndIsAFileMock

		err := dotEnv.LoadByEnv("")

		assert.Error(t, err)
	})

	t.Run("load existing env file", func(t *testing.T) {
		dotEnv := NewDotEnv()
		fakeNewEnvFile()

		mockUtils := &MockedUtils{}
		mockUtils.FileExistAndItIsAFile = mockUtils.FileExistAndIsAFileMock

		err := dotEnv.LoadByEnv("example")

		assert.NoError(t, err)

		// Clean the file
		cleanFakeFile()
	})
}

func fakeNewEnvFile() {
	file, _ := os.Create(".env.example")

	// Add fake content to it.
	_, _ = file.WriteString("TEST=123")
}

func cleanFakeFile() {
	_ = os.Remove(".env.example")
}
