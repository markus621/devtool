package utils

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/markus621/devtool/app/console"
)

//MakeDir check dir and create if not exist
func MakeDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

//DownloadFile download a file from the internet and save it to disk
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

//HTTPGet make http get request
func HTTPGet(uri string) ([]byte, error) {
	r, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}

//HTTPJson ...
func HTTPJson(uri string, v interface{}) error {
	b, err := HTTPGet(uri)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

//ExecCMD run the command in /bin/sh
func ExecCMD(dir, cmd string, env []string) (string, error) {
	c := exec.Command("/bin/sh", "-xec", fmt.Sprintln(cmd, " <&-"))
	if len(env) > 0 {
		c.Env = append(os.Environ(), env...)
	}
	if len(dir) > 0 {
		c.Dir = dir
	}
	b, err := c.CombinedOutput()
	return string(b), err
}

//WriteFile write strings to file
func WriteFile(filepath string, v ...string) error {
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

//FindInFile find text in file
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

//ValidateHash check hash
func ValidateHash(filename, hash string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return err
	}
	filehash := hex.EncodeToString(h.Sum(nil))
	if filehash != hash {
		return fmt.Errorf("invalid hash: expected[%s] actual[%s]", hash, filehash)
	}
	return nil
}

//ExtractTar ...
func ExtractTar(filename, path string) error {
	path = strings.TrimRight(path, "/") + "/"
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	rgz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	tr := tar.NewReader(rgz)
	for true {
		header, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path+header.Name, header.FileInfo().Mode().Perm()); err != nil {
				return err
			}
		case tar.TypeReg:
			console.Info("extract: %s", path+header.Name)
			newfile, err := os.OpenFile(path+header.Name, os.O_RDWR|os.O_CREATE, header.FileInfo().Mode().Perm())
			if err != nil {
				return err
			}
			if _, err := io.Copy(newfile, tr); err != nil {
				return err
			}
			newfile.Close()
		default:
			return fmt.Errorf("unknow file type")
		}
	}
	return nil
}
