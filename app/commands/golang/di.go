package golang

import (
	"github.com/deweppro/go-app/console"
)

//NewRoot root command for golang
func NewRoot() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("go", "golang")

		cmd := New()
		setter.AddCommand(cmd.GoInstall())
		setter.AddCommand(cmd.Envs())
	})
}
