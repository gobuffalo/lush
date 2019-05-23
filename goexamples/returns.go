package goexamples

import (
	"github.com/gobuffalo/lush/ast"
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

	ret := ast.NewReturned([]interface{}{"what do you get?", x})
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return &ret, nil
}
