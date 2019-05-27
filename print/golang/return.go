package golang

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astReturn(r ast.Return) error {
	var args []string

	if len(r.Statements) == 0 {
		fmt.Fprint(c, "\nreturn nil, nil")
		return nil
	}

	for _, s := range r.Statements {
		if st, ok := s.(fmt.GoStringer); ok {
			args = append(args, st.GoString())
			continue
		}
		if st, ok := s.(fmt.Stringer); ok {
			args = append(args, st.String())
		}
	}
	fmt.Fprint(c, "return golang.NewReturned(")
	fmt.Fprint(c, strings.Join(args, ", "))
	fmt.Fprintf(c, ")")
	return nil
}

func NewReturned(i ...interface{}) (*ast.Returned, error) {
	ret := ast.NewReturned(i)
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return &ret, nil
}
