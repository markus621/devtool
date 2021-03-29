package dewep

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/deweppro/go-static"
	"github.com/markus621/devtool/app/console"
	"github.com/spf13/cobra"
)

type Data struct {
	Full  string
	Midd  string
	Short string
	Dir   string
}

//NewProject dewep project generate
func NewProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "new",
		Short:   "generate project",
		Example: "devtool dewep new github.com/markus621/devtool",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("maximum number of arguments must be 1")
			}
			list := strings.Split(args[0], "/")
			if len(list) != 3 {
				return fmt.Errorf("invalid project name, must be: github.com/<USER>/<PROJECT>")
			}
			if !strings.HasPrefix(args[0], "github.com") {
				return fmt.Errorf("supported only: github.com")
			}
			return nil
		},
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		console.FatalIfErr(err, "detect path")

		params := Data{
			Full: strings.ToLower(args[0]),
			Short: func(full string) string {
				list := strings.Split(full, "/")
				return list[2]
			}(strings.ToLower(args[0])),
			Midd: func(full string) string {
				list := strings.Split(full, "/")
				return strings.Join(list[1:], "/")
			}(strings.ToLower(args[0])),
			Dir: dir,
		}

		cache := static.New()
		console.FatalIfErr(cache.FromBase64TarGZ(dewepNewProject), "unpack template")

		for _, filename := range cache.List() {
			b := cache.Get(filename)

			filename = strings.ReplaceAll(filename, "github.com/1-2-3-4-5/6-7-8-9-0", params.Full)
			filename = strings.ReplaceAll(filename, "1-2-3-4-5/6-7-8-9-0", params.Midd)
			filename = strings.ReplaceAll(filename, "6-7-8-9-0", params.Short)

			list := strings.Split(filename, "/")
			dirs := strings.Join(list[:len(list)-1], "/")

			if len(dirs) > 0 {
				console.FatalIfErr(os.MkdirAll(params.Dir+dirs, 0777), "make dir")
			}

			writer, err := os.OpenFile(params.Dir+filename, os.O_RDWR|os.O_CREATE, 0755)
			console.FatalIfErr(err, "create file")

			b = ReplaceAll(b, "github.com/1-2-3-4-5/6-7-8-9-0", params.Full)
			b = ReplaceAll(b, "1-2-3-4-5/6-7-8-9-0", params.Midd)
			b = ReplaceAll(b, "6-7-8-9-0", params.Short)

			_, err = writer.Write(b)
			console.FatalIfErr(err, "write file")
			console.FatalIfErr(writer.Close(), "close file")
		}
	}

	return cmd
}

func ReplaceAll(b []byte, old string, new string) []byte {
	return bytes.ReplaceAll(b, []byte(old), []byte(new))
}
