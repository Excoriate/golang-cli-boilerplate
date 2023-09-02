package daggerio

import (
	"context"
	"fmt"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/utils"

	"dagger.io/dagger"
)

func SetEnvVarsInContainer(c *dagger.Container, envVars map[string]string) (*dagger.Container,
	error) {
	if utils.MapIsNulOrEmpty(envVars) {
		return nil, fmt.Errorf("no environment variables are passed, skipping the environment variable configuration step")
	}

	for k, v := range envVars {
		c = c.WithEnvVariable(k, v)
	}

	return c, nil
}

// GetEnvVarsSetInContainer returns the environment variables set in the container.
func GetEnvVarsSetInContainer(ctx context.Context, c *dagger.Container) ([]dagger.EnvVariable,
	error) {
	if c == nil {
		return nil, fmt.Errorf("no container was passed")
	}

	envVars, err := c.EnvVariables(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get the environment variables from the container: %w", err)
	}

	return envVars, nil
}
