package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

func rangeExec(c *ast.Context) (*ast.Returned, error) {

	fmti, _ := c.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: c})
	fmt := fmti.(builtins.Fmt)
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

	ret := ast.NewReturned([]interface{}{myArray, myNum})
	return &ret, ret.Err()

}
