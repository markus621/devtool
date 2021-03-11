package main

import (
	"devtool/app/commands/dewep"
	"devtool/app/commands/golang"
	"devtool/app/console"

	"github.com/spf13/cobra"
)

var version = "develop"

func main() {
	root := &cobra.Command{
		Use:     "devtool",
		Short:   "develop help tool",
		Version: version,
	}

	root.AddCommand(golang.NewRoot())
	root.AddCommand(dewep.NewRoot())

	console.FatalIfErr(root.Execute(), "command execute")
}
