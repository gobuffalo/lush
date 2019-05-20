package ast

import (
	"fmt"
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
	format(a, st, verb)
}

func (a Script) MarshalAST() ([]byte, error) {
	m := map[string]interface{}{
		"Statements": a.Statements,
	}
	return toJSON("ast.Script", m)
}

func (s Script) String() string {
	return s.Statements.String()
}
