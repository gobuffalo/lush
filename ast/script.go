package ast

import (
	"fmt"
	"io"
)

type Script struct {
	Statements Statements
}

func (s Script) Exec(c *Context) (*Returned, error) {
	res, err := s.Statements.Exec(c)
	if err != nil {
		return nil, err
	}

	c.wg.Wait()

	ret, ok := res.(Returned)
	if !ok {
		return nil, nil
	}
	return &ret, nil
}

func (a Script) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		printV(st, a)
	case 's':
		io.WriteString(st, a.String())
	case 'q':
		fmt.Fprintf(st, "`%q`", a.String())
	}
}

func (a Script) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Statements": a.Statements,
	}
	return toJSON("ast.Script", m)
}

func (s Script) String() string {
	return s.Statements.String()
}
