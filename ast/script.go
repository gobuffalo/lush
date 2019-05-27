package ast

import (
	"fmt"
)

type Script struct {
	Statements Statements
}

func (s Script) Exec(c *Context) (*Returned, error) {
	res, err := s.Statements.Visit(c)
	if err != nil {
		return nil, err
	}

	c.wg.Wait()

	ret, ok := res.(Returned)
	if !ok {
		return nil, nil
	}
	return &ret, ret.Err()
}

func (a Script) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Script) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Statements": a.Statements,
	}
	return toJSON(a, m)
}

func (s Script) String() string {
	return s.Statements.String()
}
