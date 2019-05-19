package ast

import (
	"fmt"
	"io"
)

type Nil struct {
	Meta Meta
}

func (i Nil) IsZero() bool {
	return true
}

func (i Nil) String() string {
	return "nil"
}

func (i Nil) Interface() interface{} {
	return nil
}

func (i Nil) Exec(c *Context) (interface{}, error) {
	return nil, nil
}

func (i Nil) Bool(c *Context) (bool, error) {
	return false, nil
}

func (a Nil) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, a)
	case 's':
		io.WriteString(st, a.String())
	case 'q':
		fmt.Fprintf(st, "%q", a.String())
	}
}

func (a Nil) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ast.Meta": a.Meta,
	}
	return toJSON("ast.Nil", m)
}
