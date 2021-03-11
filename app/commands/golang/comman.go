package golang

import (
	"fmt"
	"os"
	"runtime"

	"devtool/app/console"
)

//noling: golint
const (
	OSLinux   = "linux"
	OSMac     = "darwin"
	OSWindows = "windows"
)

var (
	supportedOS = map[string]struct{}{
		OSLinux: {},
	}

	envFiles = map[string]interface{}{
		OSLinux: interface{}("%s/.profile"),
	}
)

//CMD model
type CMD struct {
	home   string
	osname string
	osarch string
}

//New init commands
func New() *CMD {
	cmd := &CMD{
		home:   os.Getenv("HOME"),
		osname: runtime.GOOS,
		osarch: runtime.GOARCH,
	}
	cmd.validate()
	return cmd
}

func (c *CMD) validate() {
	if _, ok := supportedOS[c.osname]; !ok {
		console.Fatal(fmt.Errorf("your OS (%s) is not supported", c.osname))
	}
}

func (c *CMD) osSwitch(list map[string]interface{}) interface{} {
	v, ok := list[c.osname]
	if !ok {
		console.Fatal(fmt.Errorf("variable is not found for your OS (%s)", c.osname))
	}
	return v
}

func (c *CMD) getEnvfile() string {
	return fmt.Sprintf(c.osSwitch(envFiles).(string), c.home)
}
