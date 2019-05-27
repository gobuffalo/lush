package ast

import (
	"fmt"
)

type Break struct {
	Meta Meta
}

func (Break) String() string {
	return "break"
}

func (a Break) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Break) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Meta": a.Meta,
	}
	return toJSON(a, m)
}

func (a Break) Visit(v Visitor) error {
	return v(a.Meta)
}
