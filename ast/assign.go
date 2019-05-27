package ast

import (
	"fmt"
	"reflect"
	"strconv"
)

type Assign struct {
	Name  Statement
	Value Statement
	Meta  Meta
}

func (a Assign) Visit(v Visitor) error {
	return v(a.Name, a.Value, a.Meta)
}

func (l Assign) String() string {
	return fmt.Sprintf("%s = %s", l.Name, l.Value)
}

func (a Assign) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Assign) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name":  a.Name,
		"Value": a.Value,
		"Meta":  a.Meta,
	}
	return toJSON(a, m)
}

func (l *Assign) Exec(c *Context) (interface{}, error) {
	if l.Value == nil {
		return nil, nil
	}

	switch t := l.Name.(type) {
	case Ident:
		name := t.String()
		if !c.Has(name) {
			return nil, l.Meta.Errorf("can not assign %s to non-existent variable", name)
		}
		si, ok := l.Value.(Execable)
		if !ok {
			c.setup(name, l.Value)
			return nil, nil
		}
		i, err := si.Exec(c)
		if err != nil {
			return nil, err
		}
		c.setup(name, i)

	case Access:
		si, ok := l.Value.(Execable)
		if !ok {
		}
		v, err := si.Exec(c)
		if err != nil {
			return nil, err
		}
		name := t.Name.String()
		if !c.Has(name) {
			return nil, l.Meta.Errorf("can not assign %s to non-existent variable", name)
		}
		k := c.Value(name)
		rv := reflect.Indirect(reflect.ValueOf(k))

		switch rv.Kind() {
		case reflect.Array, reflect.Slice:
			ind, err := strconv.Atoi(fmt.Sprint(t.Key))
			if err != nil {
				return nil, err
			}
			x := rv.Index(ind)
			x.Set(reflect.ValueOf(v))
		case reflect.Map:
			ks := fmt.Sprintf("%q", t.Key)
			if st, ok := t.Key.(fmt.Stringer); ok {
				ks = st.String()
			}
			if as, ok := t.Key.(String); ok {
				ks = as.Original
			}
			rv.SetMapIndex(reflect.ValueOf(ks), reflect.ValueOf(v))
		}
	}

	return nil, nil
}

func NewAssign(name Statement, value Statement) (*Assign, error) {
	if name.String() == "nil" {
		return nil, fmt.Errorf("can not set value for nil")
	}
	return &Assign{
		Name:  name,
		Value: value,
	}, nil
}
