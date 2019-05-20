package ast

import (
	"fmt"
	"strings"
)

type Comment struct {
	Value string
	Meta  Meta
}

func (c Comment) String() string {
	return fmt.Sprintf("// %s", c.Value)
}

func NewComment(b []byte) (Comment, error) {
	c := Comment{
		Value: string(b),
	}
	c.Value = strings.TrimSpace(c.Value)
	for _, t := range []string{"//", "#", " "} {
		c.Value = strings.TrimPrefix(c.Value, t)
	}
	return c, nil
}

func (a Comment) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Comment) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value":    genericJSON(a.Value),
		"ast.Meta": a.Meta,
	}
	return toJSON("ast.Comment", m)
}
