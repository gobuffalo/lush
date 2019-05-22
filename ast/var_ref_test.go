package ast_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func Test_VarRef(t *testing.T) {
	varRefTests := []struct {
		name string
		env  map[string]int
		in   string
		exp  int
	}{
		{
			"simple",
			map[string]int{
				"foo": 42,
			},
			"return foo",
			42,
		},
	}

	for _, test := range varRefTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewContext()
			for k, v := range test.env {
				c.Context = context.WithValue(c.Context, k, v)
			}

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
