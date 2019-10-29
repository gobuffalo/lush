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

func (c Comment) Interface() interface{} {
	return c.Value
}

func (c Comment) LushString() string {
	return c.String()
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
	switch verb {
	case 'v':
		if st.Flag('+') {
			break
		}
		if st.Flag('#') {
			break
		}
		fmt.Fprintf(st, a.Value)
		return
	}
	format(a, st, verb)
}

func (a Comment) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Value": genericJSON(a.Value),
		"Meta":  a.Meta,
	}
	return toJSON(a, m)
}
