package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

func errorsExec(c *ast.Context) (*ast.Returned, error) {

	fmti, _ := c.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: c})
	fmt := fmti.(builtins.Fmt)
	_ = fmt

	ret := ast.NewReturned([]interface{}{fmt.Errorf("stop %s", "dragging my heart around")})
	return &ret, ret.Err()

}
