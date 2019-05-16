package ast

import (
	"fmt"
	"strings"
)

type Comment struct {
	text string
	Meta Meta
}

func (c Comment) String() string {
	return fmt.Sprintf("// %s", c.text)
}

func NewComment(b []byte) (Comment, error) {
	c := Comment{
		text: string(b),
	}
	c.text = strings.TrimSpace(c.text)
	for _, t := range []string{"//", "#", " "} {
		c.text = strings.TrimPrefix(c.text, t)
	}
	return c, nil
}
