package goprint

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astBlock(b *ast.Block) error {
	if b == nil {
		return nil
	}
	if len(b.Nodes) > 0 {
		fmt.Fprintln(c)
	}

	for _, s := range b.Nodes {
		if err := blockPrint(c, s); err != nil {
			return err
		}
	}
	return nil
}

func blockPrint(c Printer, s ast.Node) error {
	// fmt.Printf("%T\n", s)
	switch t := s.(type) {
	case ast.Return:
		fmt.Fprintf(c, "return ")
		return c.astNodes(t.Nodes)
	case ast.Nodes:
		for _, st := range t {
			if err := blockPrint(c, st); err != nil {
				return err
			}
		}
	default:
		if err := c.astNode(t); err != nil {
			return err
		}
	}
	return nil
}
