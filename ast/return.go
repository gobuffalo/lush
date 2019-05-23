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
	return NewReturned(st), nil
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

func (r Return) GoString() string {
	var args []string

	if len(r.Statements) == 0 {
		return "return nil, nil"
	}

	for _, s := range r.Statements {
		if st, ok := s.(fmt.GoStringer); ok {
			args = append(args, st.GoString())
			continue
		}
		if st, ok := s.(fmt.Stringer); ok {
			args = append(args, st.String())
		}
	}

	const ret = `
ret := ast.NewReturned([]interface {}{%s})
return &ret, ret.Err()
`
	return fmt.Sprintf(ret, strings.Join(args, ", "))
}
