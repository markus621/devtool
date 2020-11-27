package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/markus621/gtool/app/utils"

	"github.com/markus621/gtool/app/console"
	"github.com/spf13/cobra"
)

const (
	goLinkTemplate = "https://golang.org/dl/go%s.linux-amd64.tar.gz"
)

var (
	regexpVersion = regexp.MustCompile("^[0-9]+\\.[0-9]+\\.[0-9]+$")
)

type GoInstall struct {
}

func (v *GoInstall) Run() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "goinstall",
		Short:   "installing golang for linux",
		Example: "gtool goinstall 1.0.0",
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		version := args[0]

		if !regexpVersion.Match([]byte(version)) {
			console.Error("invalid version: %s expects like: 1.0.0", version)
		}

		tempFile, err := ioutil.TempFile(os.TempDir(), "goinstall*.gz")
		console.FatalIfErr(err, "cant create temp tempFile")
		defer func() {
			console.FatalIfErr(os.Remove(tempFile.Name()), "cant remove temp tempFile")
		}()

		console.Progress(fmt.Sprintf("Start download version: %s", version), "---> Done", func() {
			console.FatalIfErr(utils.DownloadFile(fmt.Sprintf(goLinkTemplate, version), tempFile.Name()), "download err")
		})

		out, err := utils.ExecCMD("", "tar -C /usr/local -xzf "+tempFile.Name(), nil)
		console.Info(string(out))
		console.FatalIfErr(err, "сant unpack the archive")
	}

	return cmd
}
