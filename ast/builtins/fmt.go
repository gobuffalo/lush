package builtins

import (
	"fmt"
	"io"
)

type Setable interface {
	Set(key string, value interface{})
}

func WithFmt(c Setable, w io.Writer) {
	f := NewFmt(w)
	c.Set("errorf", f.Errorf)
	c.Set("print", f.Print)
	c.Set("printf", f.Printf)
	c.Set("println", f.Println)
	c.Set("sprint", f.Print)
	c.Set("sprintf", f.Printf)
	c.Set("sprintln", f.Println)
}

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
