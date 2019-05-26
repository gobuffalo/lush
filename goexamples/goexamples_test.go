package goexamples

import (
	"bytes"
	"fmt"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

	var rr1 interface{}
	var rr2 interface{}

	if r1 != nil {
		rr1 = *r1
	}
	if r2 != nil {
		rr2 = *r2
	}

	if rr1 == nil || rr2 == nil {
		if !(rr1 == nil && rr2 == nil) {
			fmt.Printf("! %s != %s", rr1, rr2)
			res = false
		}
	} else if !cmp.Equal(rr1, rr2, cmpopts.IgnoreUnexported(rr1, rr2)) {
		fmt.Println(cmp.Diff(rr1, rr2, cmpopts.IgnoreUnexported(rr1, rr2)))
		res = false
	}

	// if err1 == nil || err2 == nil {
	// 	return res && err1 == nil && err2 == nil
	// }
	if err1 == nil && err2 == nil {
		return res
	}

	if err1 == nil || err2 == nil {
		fmt.Printf("%s != %s", err1, err2)
		return false
	}

	if err1.Error() != err2.Error() {
		fmt.Println(cmp.Diff(err1, err2, cmpopts.IgnoreUnexported(err1, err2)))
		res = false
	}

	return res
}
