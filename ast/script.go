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

	go func() {
		for g := range c.gor {
			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				g.Exec(c)
			}()
		}
	}()
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
