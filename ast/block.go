package ast

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type Block struct {
	Nodes
	Meta Meta
}

func (b Block) String() string {
	bb := &bytes.Buffer{}
	bb.WriteString("{")
	func() {
		if len(b.Nodes) > 0 {
			if len(b.Nodes) == 1 {
				if st, ok := b.Nodes[0].(Nodes); ok {
					if len(st) == 0 {
						return
					}
				}
			}
			bb.WriteString("\n")
			x := b.Nodes.String()
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

func NewBlock(stmts ...Node) (*Block, error) {
	t := &Block{
		Nodes: Nodes(stmts),
	}
	return t, nil
}

func (a Block) Format(st fmt.State, verb rune) {
	format(a, st, verb)
}

func (a Block) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Nodes": a.Nodes,
		"Meta":  a.Meta,
	}
	return toJSON(a, m)
}
