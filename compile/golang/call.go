package golang

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Compiler) astCall(f ast.Call) error {
	if f.Concurrent {
		fmt.Fprintf(c, "go ")
	}
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
