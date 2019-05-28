package goprint

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astReturn(r ast.Return) error {
	var args []string

	if len(r.Nodes) == 0 {
		fmt.Fprint(c, "\nreturn nil, nil")
		return nil
	}

	for _, s := range r.Nodes {
		if st, ok := s.(fmt.GoStringer); ok {
			args = append(args, st.GoString())
			continue
		}
		if st, ok := s.(fmt.Stringer); ok {
			args = append(args, st.String())
		}
	}
	fmt.Fprint(c, "return runtime.Current.NewReturned(")
	fmt.Fprint(c, strings.Join(args, ", "))
	fmt.Fprintf(c, ")")
	return nil
}
