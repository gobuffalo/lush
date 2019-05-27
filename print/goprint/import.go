package goprint

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

func (c Printer) astImport(i ast.Import) error {
	x, ok := builtins.Available.Load(i.Name)
	if ok {
		s := `%si, _ := c.Imports.LoadOrStore(%q, %#v)
%s, ok := %si.(%T)
if !ok {
	return nil, fmt.Errorf("expected %T got %%T", %si)
}
_ = %s`
		fmt.Fprintf(c, s, i.Name, i.Name, x, i.Name, i.Name, x, x, i.Name, i.Name)
		fmt.Fprintf(c, "\n\n")
		return nil
	}

	s := `if _, ok := c.Imports.Load(%q); ok {
		return nil, fmt.Errorf("could not find import for %q")
	}`

	fmt.Fprintf(c, s, i.Name, i.Name)
	fmt.Fprintln(c)
	return nil
}
