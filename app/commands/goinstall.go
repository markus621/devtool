package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/markus621/devtool/app/console"
	"github.com/markus621/devtool/app/utils"
	"github.com/spf13/cobra"
)

const (
	goLinkTemplate = "https://golang.org/dl/%s.linux-amd64.tar.gz"
	goVersionList  = "https://golang.org/dl/?mode=json"
)

var (
	regexpVersion = regexp.MustCompile("^[0-9]+\\.[0-9]+\\.[0-9]+$")
)

type (
	//GoInstall model
	GoInstall struct {
	}

	//VersionResponse model
	VersionResponse struct {
		V string `json:"version"`
	}
)

//Run ...
func (v *GoInstall) Run() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "goinstall",
		Short:   "installing golang for linux",
		Example: "devtool goinstall 1.0.0\ndevtool goinstall tip",
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		version := args[0]

		if version == "tip" {
			dlv := make([]VersionResponse, 0)
			_, err := utils.HTTPRequestGET(goVersionList, &dlv)
			console.FatalIfErr(err, "cant get last version")

			if len(dlv) == 0 {
				console.Error("cant get last version")
			}
			version = dlv[0].V
		} else {
			if !regexpVersion.Match([]byte(version)) {
				console.Error("invalid version: %s expects like: 1.0.0", version)
			}
			version = "go" + version
		}

		tempFile, err := ioutil.TempFile(os.TempDir(), "goinstall*.gz")
		console.FatalIfErr(err, "cant create temp tempFile")
		defer func() {
			console.FatalIfErr(os.Remove(tempFile.Name()), "cant remove temp tempFile")
		}()

		console.Progress(fmt.Sprintf("Start download version: %s", version), "---> Done", func() {
			console.FatalIfErr(utils.DownloadFile(fmt.Sprintf(goLinkTemplate, version), tempFile.Name()), "download err")
		})

		gobinpath := utils.DevToolPath + "/go/bin"
		profile := utils.UserHomeDir() + "/.profile"

		out, err := utils.ExecCMD("", "tar -C "+utils.DevToolPath+" -xzf "+tempFile.Name(), nil)
		console.Info(string(out))
		console.FatalIfErr(err, "сant unpack the archive")

		if !utils.FindInFile(profile, gobinpath) {
			console.FatalIfErr(utils.WriteToFile(profile, "\n\nexport PATH=$PATH:"+gobinpath), "cant add install path to env")
		}
	}

	return cmd
}
