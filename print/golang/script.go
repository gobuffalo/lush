package golang

import (
	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astScript(a ast.Script) error {
	if err := c.astStatements(a.Statements); err != nil {
		return err
	}

	for _, s := range a.Statements {
		if _, ok := s.(ast.Return); ok {
			return nil
		}
	}

	return c.astReturn(ast.Return{})
}
