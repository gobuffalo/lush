package ast

import (
	"fmt"
)

type Boolable interface {
	Bool(*Context) (bool, error)
}

var True = Bool(true)
var False = Bool(false)

func NewBool(b []byte) (Bool, error) {
	bl := string(b)
	if bl == "true" {
		return True, nil
	}
	return False, nil
}

type Bool bool

func (b Bool) String() string {
	return fmt.Sprint(bool(b))
}

func (b Bool) Exec(c *Context) (interface{}, error) {
	return bool(b), nil
}

func (b Bool) Bool(c *Context) (bool, error) {
	return bool(b), nil
}

func (b Bool) Interface() interface{} {
	return bool(b)
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
