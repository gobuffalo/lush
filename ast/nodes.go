package ast

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewNodes(i interface{}) (Nodes, error) {
	var states Nodes

	ii := flatten([]interface{}{i})

	for _, i := range ii {
		switch t := i.(type) {
		case Node:
			states = append(states, t)
		case []interface{}:
			st, err := NewNodes(i)
			if err != nil {
				return states, err
			}
			states = append(states, st)
		default:
			return nil, fmt.Errorf("expected Node got %T", i)
		}
	}

	return states, nil
}

type Nodes []Node

func (t Nodes) String() string {
	var x []string
	var last Node
	for _, s := range t {
		y := strings.TrimSpace(s.String())
		if len(y) == 0 {
			continue
		}
		switch t := s.(type) {
		case Nodes:
			x = append(x, t.String())
		case Noop:
		case Comment, Import:
			x = append(x, t.String()+"\n")
		default:
			if _, ok := last.(Import); ok {
				x = append(x, "\n")
			}
			x = append(x, y+"\n\n")
		}
		last = s
	}
	return strings.Join(x, "")
}

func (t *Nodes) Append(s Node, err error) error {
	if err != nil {
		return err
	}
	(*t) = append(*t, s)
	return nil
}

func (i Nodes) Format(st fmt.State, verb rune) {
	format(i, st, verb)
}

func (st Nodes) Visit(c *Context) (interface{}, error) {
	var stmts []interface{}
	for _, s := range st {
		switch r := s.(type) {
		case Return:
			res, err := r.Visit(c)
			return res, err
		case Returned:
			return r, r.Err()
		case Break:
			return r, nil
		case Continue:
			return r, nil
		case Goroutine:
			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				r.Visit(c)
			}()
		case Visitable:
			i, err := r.Visit(c)
			if err != nil {
				return nil, err
			}
			switch t := i.(type) {
			case Returned:
				return t, t.Err()
			case Break:
				return t, nil
			case Continue:
				return t, nil
			}
			if i != nil {
				stmts = append(stmts, i)
			}
		}
	}
	return stmts, nil
}

func (st Nodes) MarshalJSON() ([]byte, error) {
	var a []interface{}
	for _, s := range st {
		a = append(a, s)
	}
	m := map[string]interface{}{
		"ast.Nodes": a,
	}

	return json.Marshal(m)
}
