package golang

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astComment(v ast.Comment) error {
	fmt.Fprintf(c, "// %s\n", v.Value)
	return nil
}
