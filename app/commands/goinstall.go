package commands

import (
	"regexp"

	"devtool/app/console"
	"devtool/app/pkg/golang"

	"github.com/spf13/cobra"
)

var (
	regexpVersion = regexp.MustCompile("^[0-9]+\\.[0-9]+\\.[0-9]+$")
)

type (
	//GoLang model
	GoLang struct {
		path string
		home string
	}
)

//GoInstall ...
func (c *CMD) GoInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "goinstall",
		Short:   "installing golang for linux",
		Example: "devtool goinstall 1.0.0\ndevtool goinstall tip",
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		version := args[0]
		var hash string

		golang := golang.New()
		upderr := golang.UpdateSettings(c.home, c.getEnvfile())

		versions, err := golang.VersionsList()
		console.FatalIfErr(err, "can't get a list of possible versions")

		if version == "tip" {
			version = versions[0].Version
		} else {
			if !regexpVersion.Match([]byte(version)) {
				console.Error("invalid version: %s expects like: 1.0.0", version)
			}
			version = "go" + version
		}

		for _, v := range versions {
			if v.Version == version {
				for _, f := range v.Files {
					if f.Version == version && f.Os == c.osname && f.Arch == c.osarch {
						version = f.Filename
						hash = f.Sha256
						break
					}
				}
			}
		}

		if len(hash) == 0 {
			console.Info("required version was not found: %s %s %s", version, c.osname, c.osarch)
		} else {
			console.Info("go: %s, %s", version, hash)
		}

		out, err := golang.Install(version, hash)
		console.FatalIfErr(err, "installation error [%s]", out)

		out, err = golang.UpdateEnv()
		console.FatalIfErr(err, "update env error [%s]", out)

		console.FatalIfErr(upderr, "update settings")
	}

	return cmd
}
