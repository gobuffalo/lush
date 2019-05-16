package ast

import (
	"fmt"
)

type Assign struct {
	Name  Ident
	Value Statement
	Meta  Meta
}

func (a *Assign) SetMeta(m Meta) {
	a.Meta = m
}

func (l Assign) String() string {
	return fmt.Sprintf("%s = %s", l.Name, l.Value)
}

func (l *Assign) Exec(c *Context) (interface{}, error) {
	if l.Value == nil {
		return nil, nil
	}

	name := l.Name.String()
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
	return nil, nil
}

func NewAssign(name Ident, value Statement) (*Assign, error) {
	if name.String() == "nil" {
		return nil, name.Meta.Errorf("can not set value for nil")
	}
	return &Assign{
		Name:  name,
		Value: value,
	}, nil
}
