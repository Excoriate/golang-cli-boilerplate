package daggerio

import (
	"fmt"

	"dagger.io/dagger"
)

func ConvertDirToDaggerDir(dir string, c *dagger.Client, options dagger.HostDirectoryOpts) (*dagger.
	Directory,
	error) {
	if dir == "" {
		return nil, fmt.Errorf("no directory was passed")
	}

	return c.Host().Directory(dir, options), nil
}

func GetMntDir() string {
	return "/mnt"
}
