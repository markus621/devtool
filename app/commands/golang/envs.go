package golang

import (
	"os"

	"github.com/deweppro/go-app/console"
)

//Envs ...
func (c *CMD) Envs() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("env", "list system environment")
		setter.Example("go env")
		setter.ExecFunc(func(args []string) {
			console.Infof("ENV:")
			for _, e := range os.Environ() {
				console.Infof(e)
			}
		})
	})
}
