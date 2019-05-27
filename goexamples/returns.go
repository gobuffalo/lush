// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/print/golang"
)

/*
x := 0

func() {
	if true {
		x = 42
	}
}()

// hopefully 42!
return "what do you get?", x
*/
func returnsExec(c *ast.Context) (*ast.Returned, error) {
	x := 0
	_ = x

	func() {
		if true {
			x = 42
		}
	}()
	// hopefully 42!
	return golang.NewReturned("what do you get?", x)
}
