package main

import (
	"github.com/markus621/devtool/app/commands"
	"github.com/markus621/devtool/app/console"
	"github.com/markus621/devtool/app/utils"
	"github.com/spf13/cobra"
)

var ver = "develop"

func main() {
	utils.Init()

	rootCMD := &cobra.Command{
		Use:     "devtool",
		Short:   "dev help tool",
		Version: ver,
	}

	rootCMD.AddCommand((&commands.GoInstall{}).Run())

	console.FatalIfErr(rootCMD.Execute(), "command execute")
}
