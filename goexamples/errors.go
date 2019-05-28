// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/gobuffalo/lush/compile/goc"
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

	return goc.NewReturned(fmt.Errorf("stop %s", "dragging my heart around"))
}
