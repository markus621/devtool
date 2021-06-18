package golang

import (
	"regexp"

	"github.com/deweppro/go-app/console"
	gopkg "github.com/markus621/devtool/app/pkg/golang"
)

var (
	regexpVersion = regexp.MustCompile(`^[0-9]+\\.[0-9]+\\.[0-9]+$`)
)

//GoInstall ...
func (c *CMD) GoInstall() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("i", "installing golang for linux")
		setter.Example("go i 1.0.0")
		setter.Example("go i tip")
		setter.Argument(1, func(s []string) ([]string, error) {
			return s, nil
		})
		setter.ExecFunc(func(args []string) {
			version := args[0]
			var hash string

			pkg := gopkg.New()
			upderr := pkg.UpdateSettings(c.home, c.getEnvfile())

			versions, err := pkg.VersionsList()
			console.FatalIfErr(err, "can't get a list of possible versions")

			if version == "tip" {
				version = versions[0].Version
			} else {
				if !regexpVersion.Match([]byte(version)) {
					console.Errorf("invalid version: %s expects like: 1.0.0", version)
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
				console.Infof("required version was not found: %s %s %s", version, c.osname, c.osarch)
			} else {
				console.Infof("go: %s, %s", version, hash)
			}

			out, err := pkg.Install(version, hash)
			console.FatalIfErr(err, "installation error [%s]", out)

			out, err = pkg.UpdateEnv()
			console.FatalIfErr(err, "update env error [%s]", out)

			console.FatalIfErr(upderr, "update settings")
		})
	})
}
