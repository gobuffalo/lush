package goprint

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astVar(v *ast.Var) error {
	if err := c.astNode(v.Name); err != nil {
		return err
	}
	fmt.Fprintf(c, " := ")
	if err := c.astNode(v.Value); err != nil {
		return err
	}
	fmt.Fprintf(c, "\t_ = ")
	if err := c.astNode(v.Name); err != nil {
		return err
	}
	fmt.Fprintf(c, "\n\n")
	return nil
}
