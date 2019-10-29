package ast

import (
	"fmt"
)

type Nil struct {
	Meta Meta
}

func (i Nil) IsZero() bool {
	return true
}

func (i Nil) String() string {
	return "nil"
}

func (i Nil) Interface() interface{} {
	return nil
}

func (i Nil) Exec(c *Context) (interface{}, error) {
	return nil, nil
}

func (i Nil) Bool(c *Context) (bool, error) {
	return false, nil
}

func (a Nil) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			break
		}
		fmt.Fprintf(st, a.String())
		return
	}
	format(a, st, verb)
}

func (a Nil) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Meta": a.Meta,
	}
	return toJSON(a, m)
}
