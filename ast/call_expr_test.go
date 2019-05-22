package ast_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

type Bar struct {
	Baz int
}

func (b Bar) Quux() int {
	return b.Baz
}

func Test_CallExpr(t *testing.T) {

	callExprTests := []struct {
		name string
		env  map[string]interface{}
		in   string
		exp  int
	}{
		{
			"single-level property access",
			map[string]interface{}{
				"foo": struct {
					Bar int
				}{42},
			},
			"return foo.Bar",
			42,
		},
		{
			"math",
			map[string]interface{}{
				"foo": struct {
					Bar int
				}{40},
			},
			"return foo.Bar + 2",
			42,
		},
		{
			"nested props",
			map[string]interface{}{
				"foo": struct {
					Bar Bar
				}{Bar{42}},
			},
			"return foo.Bar.Baz",
			42,
		},
		{
			"nested func",
			map[string]interface{}{
				"foo": struct {
					Bar Bar
				}{Bar{42}},
			},
			"return foo.Bar.Quux()",
			42,
		},
		{
			"nested func math",
			map[string]interface{}{
				"foo": struct {
					Bar Bar
				}{Bar{20}},
			},
			"return 2 * foo.Bar.Quux() + 2",
			42,
		},
	}

	for _, test := range callExprTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// set up environment
			c := NewContext()
			for k, v := range test.env {
				c.Context = context.WithValue(c.Context, k, v)
			}

			// run
			res, err := exec(test.in, c)
			if err != nil {
				t.Fatal("Err while executing: err:", err)
			}

			if !strings.EqualFold(fmt.Sprint(test.exp), fmt.Sprint(res)) {
				t.Errorf("Results differ. Got: %s, Exp: %s", fmt.Sprint(res), fmt.Sprint(test.exp))
			}

		})
	}
}
