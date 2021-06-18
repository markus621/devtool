package golang

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/deweppro/go-app/console"
	"github.com/markus621/devtool/app/utils"
)

const (
	goDirName     = ".golang"
	goDownloadURL = "https://golang.org/dl/%s"
	goVersionURL  = "https://golang.org/dl/?mode=json&include=all"

	templatedevtoolStart = `# <devtool:%s/bin>`
	templatedevtoolEnd   = `# </devtool:%s/bin>`

	templatePathEnv = `
if [ -d "%s/bin" ] ; then
    PATH="%s/bin:$PATH"
fi
`
)

//GoLang model
type GoLang struct {
	path string
	root string
}

//New golang pkg
func New() *GoLang {
	return &GoLang{
		path: os.Getenv("GOPATH"),
		root: os.Getenv("GOROOT"),
	}
}

//UpdateSettings settings install
func (c *GoLang) UpdateSettings(home, envfile string) error {
	var reboot bool
	if len(c.root) == 0 {
		c.root = fmt.Sprintf("%s/%s/go", home, goDirName)
		reboot = true
	}
	if len(c.path) == 0 {
		c.path = fmt.Sprintf("%s/%s/tools", home, goDirName)
		reboot = true
	}

	console.FatalIfErr(utils.MakeDir(c.root), "Create dir: GOROOT=%s", c.root)
	console.FatalIfErr(utils.MakeDir(c.path), "Create dir: GOPATH=%s", c.path)

	writeEnv := func(paths, key, value, filename string) bool {
		if !strings.Contains(paths, value) {
			if !utils.FindInFile(filename, fmt.Sprintf(templatedevtoolStart, value)) {
				console.FatalIfErr(
					utils.WriteFile(filename,
						fmt.Sprintf("\n"+templatedevtoolStart, value),
						fmt.Sprintf(templatePathEnv, value, value),
						fmt.Sprintf(templatedevtoolEnd+"\n", value),
					),
					"Update env file: %s", filename)
				return true
			}
		}
		return false
	}

	paths := os.Getenv("PATH")
	if writeEnv(paths, "GOROOT", c.root, envfile) {
		reboot = true
	}
	if writeEnv(paths, "GOPATH", c.path, envfile) {
		reboot = true
	}

	if reboot {
		return fmt.Errorf("settings for golang have been added to env, a reboot is required")
	}

	return nil
}

type (
	//VersionsResponse model
	VersionsResponse struct {
		Version string `json:"version"`
		Stable  bool   `json:"stable"`
		Files   []struct {
			Filename string `json:"filename"`
			Os       string `json:"os"`
			Arch     string `json:"arch"`
			Version  string `json:"version"`
			Sha256   string `json:"sha256"`
			Size     int    `json:"size"`
			Kind     string `json:"kind"`
		} `json:"files"`
	}
)

//Install custom go version
func (c *GoLang) Install(version, hash string) (string, error) {
	tmp, err := ioutil.TempFile(os.TempDir(), "goinstall*")
	if err != nil {
		return "", err
	}
	defer func() {
		if err := os.Remove(tmp.Name()); err != nil {
			console.Errorf("can`t remove temp file: %s", err.Error())
		}
	}()
	if err := utils.DownloadFile(fmt.Sprintf(goDownloadURL, version), tmp.Name()); err != nil {
		return "download error", err
	}
	if err := utils.ValidateHash(tmp.Name(), hash); err != nil {
		return "validate error", err
	}
	if err := os.RemoveAll(c.root); err != nil {
		return "can`t delete old version", err
	}
	err = utils.ExtractTar(tmp.Name(), c.root+"/../")
	return "extract error", err
}

//VersionsList get go versions
func (c *GoLang) VersionsList() (model []VersionsResponse, err error) {
	err = utils.HTTPJson(goVersionURL, &model)
	return
}

//UpdateEnv update go env
func (c *GoLang) UpdateEnv() (out string, err error) {
	out, err = utils.ExecCMD("", c.root+"/bin/go env -w GOROOT="+c.root+" GOPATH="+c.path, nil)
	return
}
