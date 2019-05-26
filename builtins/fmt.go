package builtins

import (
	"fmt"
	"io"
)

type Fmt struct {
	io.Writer
}

func NewFmt(w io.Writer) Fmt {
	return Fmt{w}
}

func (f Fmt) Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func (f Fmt) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(f, a...)
}

func (f Fmt) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(f, format, a...)
}

func (f Fmt) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(f, a...)
}

func (f Fmt) Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

func (f Fmt) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func (f Fmt) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func (f Fmt) GoString() string {
	return "builtins.Fmt{Writer: c}"
}
