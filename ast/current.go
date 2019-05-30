package ast

import "fmt"

type Current struct {
	Meta Meta
}

func (Current) Exec(c *Context) (interface{}, error) {
	return c, nil
}

func (Current) String() string {
	return "current"
}

func (c Current) Format(st fmt.State, verb rune) {
	format(c, st, verb)
}
