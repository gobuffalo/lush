package ast

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Array struct {
	Value []interface{}
	Meta  Meta
}

func (a Array) Interface() interface{} {
	return a.Value
}

func (a Array) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("[")
	var args []string
	for _, i := range a.Value {
		args = append(args, fmt.Sprint(i))
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString("]")
	return bb.String()
}

func (s Array) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, s)
	case 's':
		io.WriteString(st, s.String())
	case 'q':
		fmt.Fprintf(st, "%q", s.String())
	}
}

func (a Array) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value":    genericJSON(a.Value),
		"ast.Meta": a.Meta,
	}
	return toJSON("ast.Array", m)
}

func (a Array) Exec(c *Context) (interface{}, error) {
	var res []interface{}
	for _, i := range a.Value {
		if ex, ok := i.(Execable); ok {
			r, err := ex.Exec(c)
			if err != nil {
				return nil, err
			}
			if r != nil {
				res = append(res, r)
			}
			continue
		}
		if i != nil {
			res = append(res, i)
		}
	}
	return res, nil
}

func (a Array) Bool(c *Context) (bool, error) {
	return len(a.Value) > 0, nil
}

func NewArray(ii []interface{}) (Array, error) {
	return Array{Value: ii}, nil
}
