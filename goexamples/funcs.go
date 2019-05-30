// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/gobuffalo/lush/runtime"
)

/*
import "fmt"

f := func(a) {
	fmt.Println(a)
}

f(42)

y := func(a, b, c) {
	return (42 + 1)
}

return y(1, 2, 3)
*/
func funcsExec(current *ast.Context) (*ast.Returned, error) {
	fmti, _ := current.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: current})
	fmt, ok := fmti.(builtins.Fmt)
	if !ok {
		return nil, fmt.Errorf("expected builtins.Fmt got %T", fmti)
	}
	_ = fmt

	f := func(a interface{}) {
		fmt.Println(a)
	}
	_ = f

	f(42)
	y := func(a interface{}, b interface{}, c interface{}) interface{} {
		return (42 + 1)
	}
	_ = y

	return runtime.Current.NewReturned(y(1, 2, 3))
}
