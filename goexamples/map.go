// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/gobuffalo/lush/compile/goc"
)

/*
import "fmt"

myMap := {"a": "A", "b": 2}

fmt.Println(myMap)

for k, v := range myMap {
	fmt.Println(k)

	fmt.Println(v)
}

return myMap
*/
func mapExec(c *ast.Context) (*ast.Returned, error) {
	fmti, _ := c.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: c})
	fmt, ok := fmti.(builtins.Fmt)
	if !ok {
		return nil, fmt.Errorf("expected builtins.Fmt got %T", fmti)
	}
	_ = fmt

	myMap := map[string]interface{}{"a": "A", "b": 2}
	_ = myMap

	fmt.Println(myMap)
	for k, v := range myMap {
		_ = k
		_ = v
		fmt.Println(k)
		fmt.Println(v)
	}
	return goc.NewReturned(myMap)
}
