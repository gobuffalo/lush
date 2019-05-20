package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Float(t *testing.T) {
	table := []struct {
		in  string
		exp float64
	}{
		{`return 3.14`, 3.14},
		{`return 3.14 + 3.14`, 6.28},
		{`return 3.14 + -3.14`, 0},
		{`return 3.0 + 3`, 6.0},
		{`return 3.0 + -3`, 0},
		{`return 3.14 - -3.14`, 6.28},
		{`return 3.14 - 3.14`, 0},
		{`return 3.0 - 3`, 0.0},
		{`return 3.0 - -3`, 6.0},
		{`return 6.0 / 2.0`, 3.0},
		{`return 6.0 / 2`, 3.0},
		{`return 6.0 / -2.0`, -3.0},
		{`return 6.0 / -2`, -3.0},
		{`return 6.0 * 2.0`, 12.0},
		{`return 6.0 * 2`, 12.0},
		{`return 6.0 * -2.0`, -12.0},
		{`return 6.0 * -2`, -12.0},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			res, err := exec(tt.in, c)
			r.NoError(err)
			r.Equal(tt.exp, res.Value)

			f, err := ast.NewFloat(tt.exp)
			r.NoError(err)
			r.Equal(fmt.Sprint(tt.exp), f.String())
		})
	}
}
