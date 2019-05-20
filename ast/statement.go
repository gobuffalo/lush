package ast

import (
	"bytes"
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

func (t Statements) String() string {
	var x []string
	for _, s := range t {
		y := strings.TrimSpace(s.String())
		if len(y) == 0 {
			continue
		}
		switch t := s.(type) {
		case Statements:
			x = append(x, t.String())
		case Noop:
		case Comment:
			x = append(x, t.String()+"\n")
		default:
			x = append(x, y+"\n\n")
		}
	}
	return strings.Join(x, "")
}

func (i Statements) Format(st fmt.State, verb rune) {
	format(i, st, verb)
}

func (i Statements) MarshalJSON() ([]byte, error) {
	var st [][]byte
	for _, s := range i {
		// fmt.Printf("### ast/statement.go:70 s (%T) -> %q %+v\n", s, s, s)
		var b []byte
		var err error

		switch as := s.(type) {
		case ASTMarshaler:
			b, err = as.MarshalJSON()
		case json.Marshaler:
			b, err = as.MarshalJSON()
		default:
			b = []byte(fmt.Sprint(s))
		}

		if err != nil {
			return nil, err
		}
		st = append(st, b)
	}
	bb := &bytes.Buffer{}
	// bb.Write([]byte("["))
	res := bytes.Join(st, []byte(",\n"))
	bb.Write(res)
	// bb.Write([]byte("["))
	return bb.Bytes(), nil
}

func (st Statements) Exec(c *Context) (interface{}, error) {
	var stmts []interface{}
	for _, s := range st {
		switch r := s.(type) {
		case Return:
			res, err := r.Exec(c)
			return res, err
		case Returned:
			return r, nil
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
				return t, nil
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
