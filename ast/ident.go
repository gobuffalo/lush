package ast

import (
	"fmt"
	"reflect"
	"strings"
)

func NewIdent(b []byte) (Ident, error) {
	n := string(b)
	return Ident{Name: n}, nil
}

type Ident struct {
	Name string
	Meta Meta
}

type Idents []Ident

func (ids Idents) String() string {
	var lines []string
	for _, i := range ids {
		lines = append(lines, i.String())
	}
	return strings.Join(lines, ", ")
}

func (i Ident) IsZero() bool {
	return i == Ident{}
}

func (i Ident) String() string {
	return i.Name
}

func (a Ident) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Ident) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name": genericJSON(a.Name),
		"Meta": a.Meta,
	}
	return toJSON(a, m)
}

func (i Ident) MapKey() string {
	return i.Name
}

func (i Ident) Visit(c *Context) (interface{}, error) {
	if i.Name == "this" {
		return c, nil
	}

	if !c.Has(i.Name) {
		return nil, i.Meta.Errorf("could not find ident %s", i.Name)
	}
	return c.Value(i.Name), nil
}

func (i Ident) Bool(c *Context) (bool, error) {
	v := c.Value(i.Name)

	if v == nil {
		return false, nil
	}

	rv := reflect.Indirect(reflect.ValueOf(v))
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() > 0, nil
	case reflect.Array, reflect.Slice, reflect.Map:
		return rv.Len() > 0, nil
	}

	return true, nil
}
