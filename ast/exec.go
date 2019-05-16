package ast

import (
	"fmt"
)

type Execable interface {
	Exec(*Context) (interface{}, error)
}

type ExecStringer interface {
	Execable
	fmt.Stringer
}

func exec(c *Context, i interface{}) (interface{}, error) {
	ex, ok := i.(Execable)
	if !ok {
		return i, nil
	}
	return ex.Exec(c)
}
