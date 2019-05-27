package ast

import (
	"bytes"
	"fmt"
)

type Else struct {
	Block *Block
	Meta  Meta
}

func (a Else) Visit(v Visitor) error {
	return v(a.Block, a.Meta)
}

func (i Else) String() string {
	bb := &bytes.Buffer{}

	bb.WriteString(" else ")
	if i.Block != nil {
		bb.WriteString(i.Block.String())
	}

	return bb.String()
}

func (i Else) Bool(c *Context) (bool, error) {
	return true, nil
}

func (i Else) Exec(c *Context) (interface{}, error) {
	if i.Block == nil {
		return nil, i.Meta.Errorf("else statement missing block")
	}
	return i.Block.Exec(c)
}

func NewElse(b *Block) (Else, error) {
	return Else{
		Block: b,
	}, nil
}

func (i Else) Format(st fmt.State, verb rune) {
	format(i, st, verb)
}

func (i Else) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Block": i.Block,
		"Meta":  i.Meta,
	}
	return toJSON(i, m)
}
