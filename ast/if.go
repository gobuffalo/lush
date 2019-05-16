package ast

import (
	"bytes"
	"strings"
)

func NewIf(p Statement, e Expression, b *Block, elsa Statement) (If, error) {
	fi := If{
		PreCondition: p,
		Block:        b,
		Expression:   e,
		Clause:       elsa,
	}

	return fi, nil
}

type If struct {
	PreCondition Statement
	Expression   Expression
	Clause       Statement
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

func (i If) Bool(c *Context) (bool, error) {
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

func (i If) Exec(c *Context) (interface{}, error) {
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
