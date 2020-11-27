package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

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
