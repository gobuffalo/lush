package ast

import (
	"fmt"
	"strconv"
)

type Integer struct {
	Value int
	Meta  Meta
}

func (d Integer) Int() int {
	return d.Value
}

func (d Integer) Interface() interface{} {
	return d.Int()
}

func (d Integer) String() string {
	return strconv.Itoa(d.Value)
}

func (d Integer) Exec(c *Runtime) (interface{}, error) {
	return d.Value, nil
}

func (d Integer) Bool(c *Runtime) (bool, error) {
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
	return toJSON(a, m)
}

func (a Integer) GoString() string {
	return fmt.Sprintf("%d", a.Value)
}
