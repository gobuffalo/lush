package ast

import (
	"fmt"
)

type Continue struct {
	Meta Meta
}

func (Continue) String() string {
	return "continue"
}

func (a Continue) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Continue) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Meta": a.Meta,
	}
	return toJSON("ast.Continue", m)
}
