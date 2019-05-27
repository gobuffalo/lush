package ast

import (
	"encoding/json"
	"fmt"
	"strings"
)

type interfacer interface {
	Interface() interface{}
}

type Statement interface {
	fmt.Stringer
}

func NewStatements(i interface{}) (Statements, error) {
	var states Statements

	ii := flatten([]interface{}{i})

	for _, i := range ii {
		switch t := i.(type) {
		case Statement:
			states = append(states, t)
		case []interface{}:
			st, err := NewStatements(i)
			if err != nil {
				return states, err
			}
			states = append(states, st)
		default:
			return nil, fmt.Errorf("expected Statement got %T", i)
		}
	}

	return states, nil
}

type Statements []Statement

func (t Statements) Visit(v Visitor) error {
	for _, s := range t {
		if vs, ok := s.(Visitable); ok {
			if err := vs.Visit(v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t Statements) String() string {
	var x []string
	var last Statement
	for _, s := range t {
		y := strings.TrimSpace(s.String())
		if len(y) == 0 {
			continue
		}
		switch t := s.(type) {
		case Statements:
			x = append(x, t.String())
		case Noop:
		case Comment, Import:
			x = append(x, t.String()+"\n")
		default:
			if _, ok := last.(Import); ok {
				x = append(x, "\n")
			}
			x = append(x, y+"\n\n")
		}
		last = s
	}
	return strings.Join(x, "")
}

func (t *Statements) Append(s Statement, err error) error {
	if err != nil {
		return err
	}
	(*t) = append(*t, s)
	return nil
}

func (i Statements) Format(st fmt.State, verb rune) {
	format(i, st, verb)
}

func (st Statements) Exec(c *Context) (interface{}, error) {
	var stmts []interface{}
	for _, s := range st {
		switch r := s.(type) {
		case Return:
			res, err := r.Exec(c)
			return res, err
		case Returned:
			return r, r.Err()
		case Break:
			return r, nil
		case Continue:
			return r, nil
		case Execable:
			i, err := r.Exec(c)
			if err != nil {
				return nil, err
			}
			switch t := i.(type) {
			case Returned:
				return t, t.Err()
			case Break:
				return t, nil
			case Continue:
				return t, nil
			}
			if i != nil {
				stmts = append(stmts, i)
			}
		}
	}
	return stmts, nil
}

func (st Statements) MarshalJSON() ([]byte, error) {
	var a []interface{}
	for _, s := range st {
		a = append(a, s)
	}
	m := map[string]interface{}{
		"ast.Statements": a,
	}

	return json.Marshal(m)
}
