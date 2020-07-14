package ast

import (
	"bytes"
	"fmt"
	"strings"
)

type Return struct {
	Nodes Nodes
	Meta  Meta
}

func (r Return) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("return ")

	var lines []string
	for _, s := range r.Nodes {
		lines = append(lines, s.String())
	}
	bb.WriteString(strings.Join(lines, ", "))
	return bb.String()
}

func (r Return) Exec(c *Context) (interface{}, error) {
	st, err := r.Nodes.Exec(c)
	if err != nil {
		return NewReturned(err), err
	}
	ret := NewReturned(st)
	return ret, ret.Err()
}

func NewReturn(s Nodes) (Return, error) {
	return Return{
		Nodes: s,
	}, nil
}

func (r Return) Format(st fmt.State, verb rune) {
	format(r, st, verb)
}

func (r Return) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Nodes": r.Nodes,
		"Meta":  r.Meta,
	}
	return toJSON(r, m)
}
