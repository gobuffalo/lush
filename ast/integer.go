package ast

import (
	"fmt"
	"strconv"
)

type Integer struct {
	Value int
	Meta  Meta
}

func (d Integer) Interface() interface{} {
	return d.Value
}

func (d Integer) String() string {
	return strconv.Itoa(d.Value)
}

func (d Integer) Exec(c *Context) (interface{}, error) {
	return d.Value, nil
}

func (d Integer) Bool(c *Context) (bool, error) {
	return true, nil
}

func NewInteger(i int) (Integer, error) {
	return Integer{
		Value: i,
	}, nil
}

func (a Integer) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Integer) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value": genericJSON(a.Value),
		"Meta":  a.Meta,
	}
	return toJSON("ast.Integer", m)
}
