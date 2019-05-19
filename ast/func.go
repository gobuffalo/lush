package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func NewFunc(ax interface{}, b *Block) (Func, error) {
	var args []Ident
	for _, a := range ax.([]interface{}) {
		zz, ok := a.(Ident)
		if !ok {
			return Func{}, fmt.Errorf("expected %T to be Ident", a)
		}
		args = append(args, zz)
	}

	return Func{
		Arguments: args,
		Block:     b,
	}, nil
}

type Func struct {
	Arguments []Ident
	Block     *Block
	Meta      Meta
}

func (f Func) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("func(")
	var args []string
	for _, a := range f.Arguments {
		args = append(args, a.String())
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString(") ")
	if f.Block != nil {
		bb.WriteString(f.Block.String())
	}
	return bb.String()
}

func (f Func) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ast.Func": map[string]interface{}{
			"Arguments": f.Arguments,
			"Block":     f.Block,
			"Meta":      f.Meta,
		},
	}

	return json.MarshalIndent(m, "", "  ")
}

func (f Func) withMeta(m Meta) Statement {
	f.Meta = m
	return f
}

func (f Func) mExec(c *Context, args ...Statement) (interface{}, error) {
	c = c.Clone()
	if len(args) != len(f.Arguments) {
		return nil, f.Meta.Errorf("expected %d arguments; received %d", len(f.Arguments), len(args))
	}
	for x, i := range f.Arguments {
		v, err := exec(c, args[x])
		if err != nil {
			return nil, err
		}
		c.Set(i.Name, v)
	}
	if f.Block != nil {
		s, err := f.Block.Exec(c)
		return s, err
	}
	return nil, nil
}
