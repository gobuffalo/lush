package ast

import (
	"fmt"
	"io"
	"strconv"
)

type Integer int

func (d Integer) Interface() interface{} {
	return int(d)
}

func (d Integer) String() string {
	return strconv.Itoa(int(d))
}

func (d Integer) Exec(c *Context) (interface{}, error) {
	return int(d), nil
}

func (d Integer) Bool(c *Context) (bool, error) {
	return true, nil
}

func NewInteger(i int) (Integer, error) {
	return Integer(i), nil
}

func (a Integer) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, a)
	case 's':
		io.WriteString(st, a.String())
	case 'q':
		fmt.Fprintf(st, "%q", a.String())
	}
}

func (a Integer) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value": genericJSON(int(a)),
	}
	return toJSON("ast.Integer", m)
}
