package goexamples

import (
	"bytes"
	"fmt"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/google/go-cmp/cmp"
)

type exec func(*ast.Context) (*ast.Returned, error)

func Equal(c *ast.Context, a, b exec) bool {
	c1 := c.Clone()
	c2 := c.Clone()

	b1 := &bytes.Buffer{}
	b2 := &bytes.Buffer{}

	c1.Writer = b1
	c2.Writer = b2

	c1.Imports.Store("fmt", builtins.NewFmt(b1))
	c2.Imports.Store("fmt", builtins.NewFmt(b2))

	r1, err1 := a(c1)
	r2, err2 := b(c2)

	res := true

	if r1 != r2 {
		fmt.Println(cmp.Diff(r1, r2))
		res = false
	}

	if err1 == nil || err2 == nil {
		return res && err1 == nil && err2 == nil
	}

	if err1.Error() != err2.Error() {
		fmt.Println(cmp.Diff(r1, r2))
		res = false
	}

	return res
}
