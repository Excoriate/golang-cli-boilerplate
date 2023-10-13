package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"

	"github.com/joho/godotenv"
)

type DotEnvLoader interface {
	LoadByEnv(env string) error
	LoadByName(dotEnvFileName string) error
}

type DotEnv struct{}

func NewDotEnv() DotEnvLoader {
	return &DotEnv{}
}

func load(fileName string) error {
	if err := utils.FileExistAndItIsAFile(fileName); err != nil {
		return err
	}

	return godotenv.Load(fileName)
}

func (d *DotEnv) LoadByName(dotEnvFileName string) error {
	return load(dotEnvFileName)
}

func (d *DotEnv) LoadByEnv(env string) error {
	if env == "" {
		return fmt.Errorf("failed to load .env file, env is empty")
	}

	filename := fmt.Sprintf(".env.%s", env)
	return load(filename)
}

type MockedUtils struct {
	FileExistAndItIsAFile func(filename string) error
}

func (m *MockedUtils) FileExistAndIsAFileMock(filename string) error {
	if filename == ".env.example" {
		return nil
	}
	return fmt.Errorf("failed to load .env file, file does not exist")
}

type DotEnvFileReadOpt struct {
	FilePath       string
	IgnoreComments bool
}

func GetEnvVarsFromDotFile(opt DotEnvFileReadOpt) (map[string]string, error) {
	file, err := os.Open(opt.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	env := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" || (strings.HasPrefix(line, "#") && opt.IgnoreComments) {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // If a line does not contain '=', ignore it.
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = utils.RemoveDoubleQuotes(value)

		env[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(env) == 0 {
		return nil, fmt.Errorf("%s file is empty", opt.FilePath)
	}

	return env, nil
}
