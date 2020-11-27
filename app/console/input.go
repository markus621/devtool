package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scan  *bufio.Scanner
	yesNo = []string{"y", "n"}
)

func init() {
	scan = bufio.NewScanner(os.Stdin)
}

func printmsg(msg string, vars []string, def string) {
	defmsg := ""
	if len(def) > 0 {
		defmsg = fmt.Sprintf(" [%s]", def)
	}

	varsmsg := ""
	if len(vars) > 0 {
		varsmsg = fmt.Sprintf(" (%s)", strings.Join(vars, "/"))
	}

	Info("%s%s%s: ", msg, varsmsg, defmsg)
}

func Input(msg string, vars []string, def string) string {
	printmsg(msg, vars, def)

	for {
		if scan.Scan() {
			r := scan.Text()

			if len(r) == 0 {
				return def
			}
			if len(vars) == 0 {
				return r
			}

			for _, v := range vars {
				if v == r {
					return r
				}
			}

			printmsg("Bad answer! Try again", vars, def)
		}
	}
}

func InputBool(msg string, def bool) bool {
	sdef := "n"
	if def {
		sdef = "y"
	}

	v := Input(msg, yesNo, sdef)
	return v == "y"
}
