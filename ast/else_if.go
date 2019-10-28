package ast

import "fmt"

type ElseIf struct {
	If
}

func (e ElseIf) String() string {
	s := e.If.String()
	return " else " + s
}

func (e ElseIf) GoString() string {
	return fmt.Sprintf(" else %#v", e.If)
}

func NewElseIf(fi If) (ElseIf, error) {
	return ElseIf{
		If: fi,
	}, nil
}

func (e ElseIf) Format(st fmt.State, verb rune) {
	format(e, st, verb)
}

func (e ElseIf) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"If":   e.If,
		"Meta": e.Meta,
	}

	return toJSON(e, m)
}
