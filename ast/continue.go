package ast

import (
	"fmt"
	"io"
)

type Continue struct {
	Meta Meta
}

func (Continue) String() string {
	return "continue"
}
func (a Continue) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, a)
	case 's':
		io.WriteString(st, a.String())
	case 'q':
		fmt.Fprintf(st, "%q", a.String())
	}
}

func (a Continue) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ast.Meta": a.Meta,
	}
	return toJSON("ast.Continue", m)
}
