package ast

import (
	"bytes"
	"fmt"
	"strings"
)

func NewIf(p Node, e Expression, b *Block, elsa Node) (If, error) {
	fi := If{
		PreCondition: p,
		Block:        b,
		Expression:   e,
		Clause:       elsa,
	}

	return fi, nil
}

type If struct {
	PreCondition Node
	Expression   Expression
	Clause       Node
	Block        *Block
	Meta         Meta
}

func (i If) String() string {
	bb := &bytes.Buffer{}

	bb.WriteString("if ")
	if i.PreCondition != nil {
		bb.WriteString(strings.TrimSpace(i.PreCondition.String()) + "; ")
	}
	bb.WriteString(strings.TrimSpace(i.Expression.String()))
	bb.WriteString(" ")
	if i.Block != nil {
		bb.WriteString(i.Block.String())
	}
	if i.Clause != nil {
		bb.WriteString(i.Clause.String())
	}

	return bb.String()
}

func (i If) Format(st fmt.State, verb rune) {
	format(i, st, verb)
}

func (i If) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"PreCondition": i.PreCondition,
		"Expression":   i.Expression,
		"Clause":       i.Clause,
		"Block":        i.Block,
		"Meta":         i.Meta,
	}
	return toJSON(i, m)
}

func (i If) Bool(c *Runtime) (bool, error) {
	if epc, ok := i.PreCondition.(Expression); ok {
		b, err := epc.Bool(c)
		if err != nil {
			return false, err
		}
		if !b {
			return false, nil
		}
	}
	return i.Expression.Bool(c)
}

func (i If) Exec(c *Runtime) (interface{}, error) {
	if i.Block == nil {
		return nil, i.Meta.Errorf("if statement missing block")
	}

	if i.PreCondition != nil {
		_, err := exec(c, i.PreCondition)
		if err != nil {
			return nil, err
		}
	}

	b, err := i.Bool(c)
	if err != nil {
		return nil, err
	}
	if b {
		return i.Block.Exec(c)
	}

	if ex, ok := i.Clause.(Execable); ok {
		return ex.Exec(c)
	}
	return nil, nil
}
