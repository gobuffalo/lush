package goexamples

import (
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
)

/*
import "fmt"
// this is a comment
// there are many like it
// but this is mine
if false {
	fmt.Println("in if")
} else if (1 == 2) {
	fmt.Println("in else")
} else if true {
	fmt.Println("2 == 2")
} else {
	fmt.Println("in other else")
}


*/
func ifExec(c *ast.Context) (*ast.Returned, error) {
	fmti, _ := c.Imports.LoadOrStore("fmt", builtins.Fmt{Writer: c})
	fmt, ok := fmti.(builtins.Fmt)
	if !ok {
		return nil, fmt.Errorf("expected builtins.Fmt got %T", fmti)
	}
	_ = fmt

	// this is a comment
	// there are many like it
	// but this is mine
	if false {
		fmt.Println("in if")
	} else if 1 == 2 {
		fmt.Println("in else")
	} else if true {
		fmt.Println("2 == 2")
	} else {
		fmt.Println("in other else")
	}
	return nil, nil
}
