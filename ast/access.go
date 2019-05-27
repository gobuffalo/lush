package ast

import (
	"fmt"
	"reflect"
	"strconv"
)

func NewAccess(i Ident, key interface{}) (Access, error) {
	return Access{
		Name: i,
		Key:  key,
	}, nil
}

type Access struct {
	Name Ident
	Key  interface{}
	Meta Meta
}

func (a Access) Visit(v Visitor) error {
	return v(a.Name, a.Meta)
}

func (a Access) String() string {
	return fmt.Sprintf("%s[%v]", a.Name, a.Key)
}

func (a Access) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Access) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name": a.Name,
		"Key":  genericJSON(a.Key),
		"Meta": a.Meta,
	}
	return toJSON(a, m)
}

func (a Access) Exec(c *Context) (interface{}, error) {
	v, err := a.Name.Exec(c)
	if err != nil {
		return nil, err
	}

	rv := reflect.Indirect(reflect.ValueOf(v))
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		k := fmt.Sprint(a.Key)
		if sk, ok := a.Key.(fmt.Stringer); ok {
			k = sk.String()
		}
		i, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		if i >= rv.Len() {
			return nil, a.Meta.Errorf("index out of range %d", i)
		}
		x := rv.Index(i)
		return x.Interface(), nil
	case reflect.Map:
		k := a.Key
		if it, ok := k.(interfacer); ok {
			k = it.Interface()
		}
		x := rv.MapIndex(reflect.ValueOf(k))
		if !x.IsValid() {
			return nil, a.Meta.Errorf("could not find value for key %v", k)
		}
		return x.Interface(), nil
	}
	return nil, a.Meta.Errorf("could not access %s (%T)", a.Name, v)
}
