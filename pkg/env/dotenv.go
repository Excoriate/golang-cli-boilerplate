package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"
)

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
