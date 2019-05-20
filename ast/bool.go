package ast

import (
	"fmt"
	"io"
)

type Boolable interface {
	Bool(*Context) (bool, error)
}

var True = Bool{Value: true}
var False = Bool{Value: false}

func NewBool(b []byte) (Bool, error) {
	bl := string(b)
	if bl == "true" {
		return True, nil
	}
	return False, nil
}

type Bool struct {
	Value bool
	Meta  Meta
}

func (b Bool) String() string {
	return fmt.Sprint(b.Value)
}

func (b Bool) Exec(c *Context) (interface{}, error) {
	return b.Value, nil
}

func (b Bool) Bool(c *Context) (bool, error) {
	return b.Value, nil
}

func (b Bool) Interface() interface{} {
	return b.Value
}

func (a Bool) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, a)
	case 's':
		io.WriteString(st, a.String())
	case 'q':
		fmt.Fprintf(st, "%q", a.String())
	}
}

func (a Bool) MarshalAST() ([]byte, error) {
	m := map[string]interface{}{
		"Value": genericJSON(a.Value),
	}
	return toJSON("ast.Bool", m)
}

func boolExec(s interface{}, c *Context) (bool, error) {
	if b, ok := s.(Boolable); ok {
		return b.Bool(c)
	}
	if b, ok := s.(bool); ok {
		return b, nil
	}
	return false, nil
}
