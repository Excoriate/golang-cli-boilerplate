package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileExistAndItIsAFile(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		f, err := ioutil.TempFile("", "")
		require.NoError(t, err)
		_, err = f.WriteString("test") // Write something to file
		require.NoError(t, err)
		require.NoError(t, f.Close())
		assert.NoError(t, FileExistAndItIsAFile(f.Name()))
		_ = os.Remove(f.Name())
	})

	t.Run("non-existent file", func(t *testing.T) {
		assert.Error(t, FileExistAndItIsAFile("non_existent.txt"))
	})
}

func TestWriteFileWithContent(t *testing.T) {
	t.Run("write valid file", func(t *testing.T) {
		filename := "valid_file.txt"
		require.NoError(t, WriteFileWithContent(filename, []byte("Hello, World!")))
		assert.NoError(t, FileExistAndItIsAFile(filename))
		_ = os.Remove(filename)
	})

	t.Run("write to invalid filepath", func(t *testing.T) {
		require.Error(t, WriteFileWithContent("/invalid/invalid_file.txt", []byte("Hello, World!")))
	})

	t.Run("empty filename", func(t *testing.T) {
		require.Error(t, WriteFileWithContent("", []byte("Hello, World!")))
	})

	t.Run("empty content", func(t *testing.T) {
		filename := "valid_file.txt"
		require.Error(t, WriteFileWithContent(filename, []byte("")))
		assert.Error(t, FileExistAndItIsAFile(filename))
		_ = os.Remove(filename)
	})
}
