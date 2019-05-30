// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

/*
import "fmt"

fmt.Println(current)
*/
func currentExec(current *ast.Context) (*ast.Returned, error) {
	fmti, _ := current.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: current})
	fmt, ok := fmti.(builtins.Fmt)
	if !ok {
		return nil, fmt.Errorf("expected builtins.Fmt got %T", fmti)
	}
	_ = fmt

	fmt.Println(current)

	return nil, nil
}