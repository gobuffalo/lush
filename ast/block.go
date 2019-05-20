package ast

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type Block struct {
	Statements
	Meta Meta
}

func (b Block) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("{")
	if len(b.Statements) > 0 {
		bb.WriteString("\n")
		x := b.Statements.String()
		x = strings.TrimSpace(x)
		scan := bufio.NewScanner(strings.NewReader(x))
		for scan.Scan() {
			s := scan.Text()
			if len(strings.TrimSpace(s)) == 0 {
				bb.WriteString("\n")
				continue
			}
			bb.WriteString(fmt.Sprintf("\t%s\n", s))
		}
	}
	bb.WriteString("}")
	return bb.String()
}

func NewBlock(stmts Statements) (*Block, error) {
	t := &Block{
		Statements: stmts,
	}
	return t, nil
}

func (a Block) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Block) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"ast.Statements": a.Statements,
		"ast.Meta":       a.Meta,
	}
	return toJSON("ast.Block", m)
}
