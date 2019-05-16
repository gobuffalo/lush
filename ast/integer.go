package ast

import (
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
