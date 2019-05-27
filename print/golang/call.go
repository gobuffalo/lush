package golang

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astCall(f ast.Call) error {
	fmt.Fprintf(c, f.Name.String())
	if (f.FName != ast.Ident{}) {
		fmt.Fprintf(c, ".")
		fmt.Fprintf(c, f.FName.String())
	}
	fmt.Fprintf(c, "(")
	var args []string
	for _, a := range f.Arguments {
		st := a.(fmt.Stringer)
		args = append(args, strings.TrimSpace(st.String()))
	}
	fmt.Fprintf(c, strings.Join(args, ", "))
	fmt.Fprintf(c, ")\n")
	return nil
}

func (c Printer) astGoroutine(g ast.Goroutine) error {
	fmt.Fprintf(c, "go %s\n", g.Call)
	return nil
}
