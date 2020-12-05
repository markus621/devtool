package console

import (
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// Warn printf yellow
func Warn(msg string, args ...interface{}) {
	color.Yellow(msg, args...)
}

// Fatal print red and os.exit
func Fatal(err error) {
	color.Red("[ERR] %s", err.Error())
	os.Exit(1)
}

// Error print red and os.exit
func Error(msg string, args ...interface{}) {
	color.Red(msg, args...)
	os.Exit(1)
}

// Info prints white
func Info(msg string, args ...interface{}) {
	color.White(msg, args...)
}

// Debug prints hiwhite
func Debug(msg string, args ...interface{}) {
	color.HiWhite(msg, args...)
}

// FatalIfErr calls fatal if err not nil
func FatalIfErr(err error, comment string) {
	if err != nil {
		Fatal(errors.Wrap(err, comment))
	}
}

//Progress console progress log
func Progress(start, finish string, call func()) {
	Info(start)
	call()
	Info(finish)
}
