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
	func() {
		if len(b.Statements) > 0 {
			if len(b.Statements) == 1 {
				if st, ok := b.Statements[0].(Statements); ok {
					if len(st) == 0 {
						return
					}
				}
			}
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
	}()
	bb.WriteString("}")
	return bb.String()
}

func NewBlock(stmts ...Statement) (*Block, error) {
	t := &Block{
		Statements: Statements(stmts),
	}
	return t, nil
}

func (a Block) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Block) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Statements": a.Statements,
		"Meta":       a.Meta,
	}
	return toJSON(a, m)
}

func (b Block) GoString() string {
	bb := &bytes.Buffer{}
	if len(b.Statements) > 0 {
		bb.WriteString("\n")
		x := b.Statements.GoString()
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
	return bb.String()
}
