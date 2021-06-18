package dewep

import (
	"github.com/deweppro/go-app/console"
)

var dewepNewProject = ""

//NewRoot root command for golang
func NewRoot() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("dewep", "helper for github.com/deweppro")
		setter.AddCommand(NewService())
		setter.AddCommand(NewProject())
	})
}
