// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/gobuffalo/lush/runtime"
)

/*
import "fmt"

myNum := 0

myArray := [1, "2", true]

for i, x := range myArray {
	myNum = i

	fmt.Print(i, x)
}

return myArray, myNum
*/
func rangeExec(c *ast.Runtime) (*ast.Returned, error) {
	fmti, _ := c.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: c})
	fmt, ok := fmti.(builtins.Fmt)
	if !ok {
		return nil, fmt.Errorf("expected builtins.Fmt got %T", fmti)
	}
	_ = fmt

	myNum := 0
	_ = myNum

	myArray := []interface{}{1, "2", true}
	_ = myArray

	for i, x := range myArray {
		_ = i
		_ = x
		myNum = i
		fmt.Print(i, x)
	}
	return runtime.Current.NewReturned(myArray, myNum)
}
