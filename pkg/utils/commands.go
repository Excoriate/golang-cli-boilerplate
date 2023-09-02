package utils

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/google/shlex"
)

type Command struct {
	Binary   string
	Commands []string
}

// GetCommandArgs parses the job command and returns the arguments.
func GetCommandArgs(cmd string) ([]string, error) {
	args, err := shlex.Split(cmd)

	if err != nil {
		return nil, fmt.Errorf("could not parse the command: %s", err)
	}

	return args, nil
}

// RunCommand runs a command and returns the output.
func RunCommand(cmd Command) (string, error) {
	var c *exec.Cmd
	switch {
	case cmd.Binary != "":
		c = exec.Command(cmd.Binary,
			cmd.Commands...)
	case len(cmd.Commands) > 0:
		c = exec.Command(cmd.Commands[0],
			cmd.Commands[1:]...)
	default:
		return "", fmt.Errorf("no command provided")
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr

	err := c.Run()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v, stderr: %s", err, stderr.String())
	}

	return out.String(), nil
}
