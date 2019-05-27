package goprint

import (
	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astScript(a ast.Script) error {
	if err := c.astNodes(a.Nodes); err != nil {
		return err
	}

	for _, s := range a.Nodes {
		if _, ok := s.(ast.Return); ok {
			return nil
		}
	}

	return c.astReturn(ast.Return{})
}
