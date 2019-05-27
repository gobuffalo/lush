package golang

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astVar(v *ast.Var) error {
	if err := c.astStatement(v.Name); err != nil {
		return err
	}
	fmt.Fprintf(c, " := ")
	if err := c.astStatement(v.Value); err != nil {
		return err
	}
	fmt.Fprintf(c, "\t_ = ")
	if err := c.astStatement(v.Name); err != nil {
		return err
	}
	fmt.Fprintf(c, "\n\n")
	return nil
}
