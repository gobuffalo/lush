package ast

import (
	"fmt"
	"io"
)

type Break struct {
	Meta Meta
}

func (Break) String() string {
	return "break"
}

func (a Break) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, a)
	case 's':
		io.WriteString(st, a.String())
	case 'q':
		fmt.Fprintf(st, "%q", a.String())
	}
}

func (a Break) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ast.Meta": a.Meta,
	}
	return toJSON("ast.Break", m)
}
