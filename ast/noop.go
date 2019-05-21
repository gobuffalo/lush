package ast

import "fmt"

type Noop struct {
	Text string
	Meta Meta
}

func (n Noop) String() string {
	return ""
}

func NewNoop(b []byte) (Noop, error) {
	return Noop{
		Text: string(b),
	}, nil
}

func (n Noop) Format(st fmt.State, verb rune) {
	format(n, st, verb)
}

func (n Noop) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Text": n.Text,
		"Meta": n.Meta,
	}
	return toJSON("ast.Noop", m)
}
