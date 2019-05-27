package ast

import (
	"bytes"
	"fmt"
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

func (a Range) Visit(v Visitor) error {
	return v(a.For, a.Meta)
}

func (f Range) Format(st fmt.State, verb rune) {
	format(f, st, verb)
}

func (f Range) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"For":  f.For,
		"Meta": f.Meta,
	}

	return toJSON(f, m)
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

func (f Range) GoString() string {
	bb := &bytes.Buffer{}
	bb.WriteString("for ")
	var args []string
	for _, a := range f.Args {
		args = append(args, a.String())
	}
	bb.WriteString(strings.Join(args, ", "))
	bb.WriteString(" := range ")
	bb.WriteString(f.Name.String() + " {\n")
	for _, a := range f.Args {
		fmt.Fprintf(bb, "\n\t_ = %s\n", a)
	}
	if f.Block != nil {
		bb.WriteString(f.Block.GoString())
	}

	bb.WriteString("}")
	return bb.String()
}
