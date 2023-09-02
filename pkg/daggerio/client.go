package daggerio

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func NewClient(ctx context.Context) (*dagger.
	Client, error) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to dagger client ("+
			"no rootDir passed to be mounted): %w", err)
	}

	return client, nil
}
