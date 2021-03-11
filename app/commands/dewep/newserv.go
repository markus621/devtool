package dewep

import (
	"devtool/app/console"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const tmpl = `
package {{.Package}}

//{{.Name}} service model
type {{.Name}} struct {
	
}

//New{{.Name}} init service
func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{}
}

//Up start service
func (o *{{.Name}}) Up() error {
	return nil
}

//Down stop service
func (o *{{.Name}}) Down() error {
	return nil
}

`

//Tmpl model
type Tmpl struct {
	Name     string
	Package  string
	Filename string
}

//NewService dewep service generate
func NewService() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serv",
		Short:   "generate service",
		Example: "devtool dewep serv Hello",
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		v := &Tmpl{Name: args[0]}
		v.Filename = strings.ToLower(v.Name) + "_service.go"
		dir, err := os.Getwd()
		console.FatalIfErr(err, "detect path")
		v.Package = filepath.Base(dir)

		parse, err := template.New("tmpl").Parse(tmpl)
		console.FatalIfErr(err, "template decode")

		r, err := os.OpenFile(v.Filename, os.O_RDWR|os.O_CREATE, 0755)
		console.FatalIfErr(err, "create file %s", v.Filename)
		defer r.Close()

		console.FatalIfErr(parse.Execute(r, v), "template generate")

		console.Info("Done")
	}

	return cmd
}
