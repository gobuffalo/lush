package ast

import (
	"fmt"
)

type Float struct {
	Value float64
	Meta  Meta
}

func (d Float) Interface() interface{} {
	return d.Float()
}

func (d Float) Float() float64 {
	return d.Value
}

func (n Float) String() string {
	return fmt.Sprint(n.Value)
}

func (n Float) Visit(c *Context) (interface{}, error) {
	return n.Value, nil
}

func NewFloat(f float64) (Float, error) {
	ft := Float{
		Value: f,
	}
	return ft, nil
}

func (n Float) Format(st fmt.State, verb rune) {
	format(n, st, verb)
}

func (n Float) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value": n.Value,
		"Meta":  n.Meta,
	}
	return toJSON(n, m)
}
