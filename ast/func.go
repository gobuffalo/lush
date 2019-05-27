package ast

import (
	"bytes"
	"fmt"
	"strings"
)

func NewFunc(ax interface{}, b *Block) (Func, error) {
	var args Idents
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
	Arguments Idents
	Block     *Block
	Meta      Meta
}

func (f Func) Format(st fmt.State, verb rune) {
	format(f, st, verb)
}

func (f Func) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Arguments": f.Arguments,
		"Block":     f.Block,
		"Meta":      f.Meta,
	}

	return toJSON(f, m)
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

func (f Func) mExec(c *Context, args ...Node) (interface{}, error) {
	c = c.Clone()
	if len(args) != len(f.Arguments) {
		return nil, f.Meta.Errorf("expected %d arguments; received %d", len(f.Arguments), len(args))
	}
	for x, i := range f.Arguments {
		c.Set(i.Name, args[x])
	}
	if f.Block != nil {
		s, err := f.Block.Visit(c)
		return s, err
	}
	return nil, nil
}
