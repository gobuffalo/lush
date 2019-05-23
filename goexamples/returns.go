package goexamples

import (
	"github.com/gobuffalo/lush/ast"
)

func returnsExec(c *ast.Context) (*ast.Returned, error) {

	x := 0
	_ = x

	func() {
		if true {
			x = 42
		}
	}()
	// hopefully 42!

	ret := ast.NewReturned([]interface{}{"what do you get?", x})
	return &ret, ret.Err()

}
