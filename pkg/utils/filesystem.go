package utils

import (
	"fmt"
	"os"
)

func FileExistAndItIsAFile(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty file path")
	}

	currentDir, _ := os.Getwd()

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist in current directory %s", filePath, currentDir)
		}

		return fmt.Errorf("error checking the file %s: %v", filePath, err)
	}

	return nil
}

func FileIsNotEmpty(filepath string) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", file, err)
	}

	if len(file) == 0 {
		return fmt.Errorf("file %s is empty", filepath)
	}

	return nil
}

func DirExistAndHasContent(dirPath string) error {
	if dirPath == "" {
		return fmt.Errorf("empty directory path")
	}

	currentDir, _ := os.Getwd()

	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory %s does not exist in current directory %s", dirPath, currentDir)
		}

		return fmt.Errorf("error checking the directory %s: %v", dirPath, err)
	}

	return nil
}

func WriteFileWithContent(filename string, content []byte) error {
	if filename == "" {
		return fmt.Errorf("failed to write a file, filename is empty")
	}

	if len(content) == 0 {
		return fmt.Errorf("failed to write a file, content is empty")
	}

	err := os.WriteFile(filename, content, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write a file, error: %s", err.Error())
	}

	return nil
}

func GetHomeDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}

func GetCurrentDir() string {
	currentDir, _ := os.Getwd()
	return currentDir
}
