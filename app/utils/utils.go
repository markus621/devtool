package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var (
	//DevToolPath ...
	DevToolPath string
	//UserHomePath ...
	UserHomePath string
)

//Init ...
func Init() {
	UserHomePath = UserHomeDir()
	DevToolPath = UserHomePath + "/.devtool"

	if _, err := os.Stat(DevToolPath); os.IsNotExist(err) {
		os.Mkdir(DevToolPath, 0744)
	}
}

//DownloadFile ...
func DownloadFile(uri, filepath string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

//HTTPRequestGET ...
func HTTPRequestGET(uri string, v interface{}) ([]byte, error) {
	r, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if v != nil {
		err = json.Unmarshal(b, v)
	}
	return b, err
}

//ExecCMD ...
func ExecCMD(dir, cmd string, env []string) ([]byte, error) {
	c := exec.Command("/bin/sh", "-xec", fmt.Sprintln(cmd, " <&-"))

	if len(env) > 0 {
		c.Env = append(os.Environ(), env...)
	}
	if len(dir) > 0 {
		c.Dir = dir
	}

	return c.CombinedOutput()
}

//FindInFile ...
func FindInFile(filepath, v string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), v) {
			return true
		}
	}

	return false
}

//UserHomeDir ...
func UserHomeDir() string {
	u, err := user.Current()
	if err != nil {
		return "."
	}

	return u.HomeDir
}

//WriteToFile ...
func WriteToFile(filepath string, v ...string) error {
	var (
		f   *os.File
		fi  os.FileInfo
		err error
	)

	fi, err = os.Lstat(filepath)
	if err == nil {
		f, err = os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, fi.Mode().Perm())
	} else {
		f, err = os.Create(filepath)
	}

	if err != nil {
		return err
	}

	defer f.Close()

	for _, vi := range v {
		if _, err = f.WriteString(vi); err != nil {
			return err
		}
	}
	return nil
}
