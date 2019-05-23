package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

/*
import "fmt"

return fmt.Errorf("stop %s", "dragging my heart around")
*/
func errorsExec(c *ast.Context) (*ast.Returned, error) {

	fmti, _ := c.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: c})
	fmt, ok := fmti.(builtins.Fmt)
	if !ok {
		return nil, fmt.Errorf("expected builtins.Fmt got %T", fmti)
	}
	_ = fmt

	ret := ast.NewReturned([]interface{}{fmt.Errorf("stop %s", "dragging my heart around")})
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return &ret, nil

}
