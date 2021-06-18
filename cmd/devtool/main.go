package main

import (
	"github.com/deweppro/go-app/console"
	"github.com/markus621/devtool/app/commands/dewep"
	"github.com/markus621/devtool/app/commands/golang"
)

func main() {
	root := console.New("devtool", "develop help tool")
	root.AddCommand(golang.NewRoot())
	root.AddCommand(dewep.NewRoot())
	root.Exec()
}
