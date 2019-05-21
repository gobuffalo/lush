package ast

import (
	"bytes"
	"fmt"
	"strings"
)

type Return struct {
	Statements Statements
	Meta       Meta
}

func (r Return) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("return ")
	var lines []string
	for _, s := range r.Statements {
		lines = append(lines, s.String())
	}
	bb.WriteString(strings.Join(lines, ", "))
	return bb.String()
}

func (r Return) Exec(c *Context) (interface{}, error) {
	st, err := r.Statements.Exec(c)
	if err != nil {
		return nil, err
	}
	return NewReturned(st), err
}

func NewReturn(s Statements) (Return, error) {
	return Return{
		Statements: s,
	}, nil
}

func (r Return) Format(st fmt.State, verb rune) {
	format(r, st, verb)
}

func (r Return) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Statements": r.Statements,
		"Meta":       r.Meta,
	}
	return toJSON(r, m)
}
