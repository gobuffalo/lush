package ast

import (
	"fmt"
	"strings"
)

type Shebang struct {
	Value string
	Meta  Meta
}

func (c Shebang) String() string {
	return c.Value
}

func NewShebang(b []byte) (Shebang, error) {
	c := Shebang{
		Value: string(b),
	}
	c.Value = strings.TrimSpace(c.Value)
	return c, nil
}

func (a Shebang) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Shebang) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value": genericJSON(a.Value),
		"Meta":  a.Meta,
	}
	return toJSON(a, m)
}
