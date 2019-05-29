package ast

import (
	"fmt"
)

type Script struct {
	Nodes Nodes
}

func (s Script) Exec(c *Context) (*Returned, error) {
	res, err := s.Nodes.Exec(c)
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
		"Nodes": a.Nodes,
	}
	return toJSON(a, m)
}

func (s Script) String() string {
	return s.Nodes.String()
}
