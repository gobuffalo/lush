package ast

import (
	"bytes"
	"fmt"
	"strings"
)

type Array struct {
	Value []interface{}
	Meta  Meta
}

func (a Array) Interface() interface{} {
	return a.Slice()
}

func (a Array) Slice() []interface{} {
	return a.Value
}

func (a Array) GoString() string {
	if a.Value == nil {
		a.Value = []interface{}{}
	}
	return fmt.Sprintf("%#v", a.Value)
}

func (a Array) LushString() string {
	return a.String()
}

func (a Array) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("[")
	var args []string
	for _, i := range a.Value {
		if s, ok := i.(fmt.Stringer); ok {
			args = append(args, s.String())
			continue
		}
		args = append(args, fmt.Sprint(i))
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString("]")
	return bb.String()
}

func (s Array) Format(st fmt.State, verb rune) {
	format(s, st, verb)
}

func (a Array) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value": genericJSON(a.Value),
		"Meta":  a.Meta,
	}
	return toJSON(a, m)
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

func (a Array) Len() int {
	return len(a.Value)
}

func (a Array) Less(i, j int) bool {
	return fmt.Sprint(a.Value[i]) < fmt.Sprint(a.Value[j])
}

func (a Array) Swap(i, j int) {
	a.Value[i], a.Value[j] = a.Value[j], a.Value[i]
}
