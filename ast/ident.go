package ast

import (
	"reflect"
)

func NewIdent(b []byte) (Ident, error) {
	n := string(b)
	return Ident{Name: n}, nil
}

type Ident struct {
	Name string
	Meta Meta
}

func (i Ident) IsZero() bool {
	return i == Ident{}
}

func (i Ident) String() string {
	return i.Name
}

func (i Ident) MapKey() string {
	return i.Name
}

func (i Ident) Exec(c *Context) (interface{}, error) {
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
