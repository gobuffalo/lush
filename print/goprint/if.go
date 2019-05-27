package goprint

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astIf(i ast.If) error {

	fmt.Fprintf(c, "if ")
	if i.PreCondition != nil {
		fmt.Fprintf(c, strings.TrimSpace(i.PreCondition.String())+"; ")
	}
	fmt.Fprintf(c, strings.TrimSpace(i.Expression.String()))
	fmt.Fprintf(c, " ")
	if i.Block != nil {
		fmt.Fprintf(c, i.Block.String())
	}
	if i.Clause != nil {
		fmt.Fprintf(c, i.Clause.String())
	}

	return nil
}
