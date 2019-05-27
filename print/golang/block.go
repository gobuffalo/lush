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
	return c.astStatements(b.Statements)
}
