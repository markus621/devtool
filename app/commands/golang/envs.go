package golang

import (
	"os"

	"github.com/markus621/devtool/app/console"
	"github.com/spf13/cobra"
)

//Envs ...
func (c *CMD) Envs() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "env",
		Short:   "list system environment",
		Example: "devtool go env",
		Args:    cobra.MinimumNArgs(0),
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		console.Info("ENV:")
		for _, e := range os.Environ() {
			console.Info(e)
		}
	}

	return cmd
}
