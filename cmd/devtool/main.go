package main

import (
	"devtool/app/commands"
	"devtool/app/console"

	"github.com/spf13/cobra"
)

var version = "develop"

func main() {
	rootCMD := &cobra.Command{
		Use:     "devtool",
		Short:   "dev help tool",
		Version: version,
	}

	cmd := commands.New()

	rootCMD.AddCommand(cmd.GoInstall())
	rootCMD.AddCommand(cmd.Envs())

	console.FatalIfErr(rootCMD.Execute(), "command execute")
}
