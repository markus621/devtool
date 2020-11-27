package main

import (
	"github.com/markus621/gtool/app/commands"
	"github.com/markus621/gtool/app/console"
	"github.com/spf13/cobra"
)

var Version = "develop"

func main() {
	rootCMD := &cobra.Command{
		Use:     "gtool",
		Short:   "golang help tool",
		Version: Version,
	}

	rootCMD.AddCommand((&commands.GoInstall{}).Run())

	console.FatalIfErr(rootCMD.Execute(), "command execute")
}
