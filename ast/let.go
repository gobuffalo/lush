package ast

import (
	"fmt"
)

func NewLet(name Ident, value Statement) (*Let, error) {
	if name.String() == "nil" {
		return nil, name.Meta.Errorf("can not set value for nil")
	}
	return &Let{
		Name:  name,
		Value: value,
	}, nil
}

type Let struct {
	Name  Ident
	Value Statement
	Meta  Meta
}

func (b *Let) SetMeta(m Meta) {
	b.Meta = m
}

func (l Let) String() string {
	return fmt.Sprintf("let %s = %s", l.Name, l.Value)
}

func (l *Let) Exec(c *Context) (interface{}, error) {
	if l.Value == nil {
		return nil, nil
	}
	si, ok := l.Value.(Execable)
	if !ok {
		c.Set(l.Name.String(), l.Value)
		return nil, nil
	}
	i, err := si.Exec(c)
	if err != nil {
		return nil, err
	}
	c.Set(l.Name.String(), i)
	return nil, nil
}
