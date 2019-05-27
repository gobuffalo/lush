package golang

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astBlock(b *ast.Block) error {
	if b == nil {
		return nil
	}
	if len(b.Statements) > 0 {
		fmt.Fprintln(c)
	}

	for _, s := range b.Statements {
		if err := blockPrint(c, s); err != nil {
			return err
		}
	}
	return nil
}

func blockPrint(c Printer, s ast.Statement) error {
	// fmt.Printf("%T\n", s)
	switch t := s.(type) {
	case ast.Return:
		fmt.Fprintf(c, "return ")
		return c.astStatements(t.Statements)
	case ast.Statements:
		for _, st := range t {
			if err := blockPrint(c, st); err != nil {
				return err
			}
		}
	default:
		if err := c.astStatement(t); err != nil {
			return err
		}
	}
	return nil
}
