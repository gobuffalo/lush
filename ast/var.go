package ast

import (
	"errors"
	"fmt"
)

func NewVar(name Ident, value Statement) (*Var, error) {
	if name.String() == "nil" {
		return nil, errors.New("can not set value for nil")
	}
	return &Var{
		Name:  name,
		Value: value,
	}, nil
}

type Var struct {
	Name  Ident
	Value Statement
	Meta  Meta
}

func (l Var) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Name":     l.Name,
		"Value":    l.Value,
		"ast.Meta": l.Meta,
	}
	return toJSON("ast.Var", m)
}

func (l Var) String() string {
	return fmt.Sprintf("%s := %s", l.Name, l.Value)
}

func (l *Var) Exec(c *Context) (interface{}, error) {
	if l.Value == nil {
		return nil, nil
	}
	name := l.Name.String()
	if c.Has(name) {
		return nil, l.Meta.Errorf("can not assign %s to existing variable", name)
	}
	si, ok := l.Value.(Execable)
	if !ok {
		c.Set(name, l.Value)
		return nil, nil
	}
	i, err := si.Exec(c)
	if err != nil {
		return nil, l.Meta.Wrap(err)
	}
	c.Set(name, i)
	return nil, nil
}

func (l Var) Format(st fmt.State, verb rune) {
	format(l, st, verb)
}
