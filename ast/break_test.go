package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Break_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `break`},
		{`%q`, `"break"`},
		{`%v`, `break`},
		{`%+v`, `break`},
		{`%#v`, `break`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, ast.Break{})

			r.Equal(tt.out, ft)
		})
	}
}
