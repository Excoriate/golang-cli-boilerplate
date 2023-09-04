package cliutils

import "github.com/spf13/cobra"

// GetCMDContext returns the context of a cobra command.
func GetCMDContext(c *cobra.Command, key string) interface{} {
	ctx := c.Context().Value(key)
	return ctx
}
