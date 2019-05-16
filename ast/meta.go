package ast

import (
	"fmt"
)

type Meta struct {
	Filename string
	Line     int
	Col      int
	Offset   int
	Original string
}

func (m Meta) String() string {
	if len(m.Filename) == 0 {
		return fmt.Sprintf("%d:%d", m.Line, m.Col)
	}
	return fmt.Sprintf("%s %d:%d", m.Filename, m.Line, m.Col)
}

func (m Meta) Wrap(err error) error {
	return fmt.Errorf("%s: %s", m.String(), err)
}

func (m Meta) Errorf(format string, args ...interface{}) error {
	return m.Wrap(fmt.Errorf(format, args...))
}
