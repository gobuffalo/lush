package golang

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astRange(f ast.Range) error {
	fmt.Fprintf(c, "for ")
	var args []string
	for _, a := range f.Args {
		args = append(args, a.String())
	}
	fmt.Fprintf(c, strings.Join(args, ", "))
	fmt.Fprintf(c, " := range ")
	fmt.Fprintf(c, f.Name.String()+" {")
	for _, a := range f.Args {
		fmt.Fprintf(c, "\n\t_ = %s", a)
	}
	if f.Block != nil {
		if err := c.astBlock(f.Block); err != nil {
			return err
		}
	}

	fmt.Fprintf(c, "}\n")
	return nil
}
