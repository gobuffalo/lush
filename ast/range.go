package ast

import (
	"bytes"
	"strings"
)

func NewRange(n ExecStringer, args interface{}, b *Block) (Range, error) {
	r := Range{}
	f, err := NewFor(n, args, b)
	if err != nil {
		return r, err
	}
	f.normalSingle = true
	r.For = f
	return r, nil
}

type Range struct {
	For
	Meta Meta
}

func (f Range) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("for ")
	var args []string
	for _, a := range f.Args {
		args = append(args, a.String())
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString(" := range ")
	bb.WriteString(f.Name.String() + " ")
	if f.Block != nil {
		bb.WriteString(f.Block.String())
	}
	return bb.String()
}
