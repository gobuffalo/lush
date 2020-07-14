package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_Integer(t *testing.T) {
	table := []struct {
		in  string
		exp float64
	}{
		{`return 314`, 314},
		{`return 314 + 314`, 628},
		{`return 314 + -314`, 0},
		{`return 30 + 3`, 33},
		{`return 30 + -3`, 27},
		{`return 314 - -314`, 628},
		{`return 314 - 314`, 0},
		{`return 30 - 3`, 27},
		{`return 30 - -3`, 33},
		{`return 60 / 20`, 3.0},
		{`return 60 / -20`, -3.0},
		{`return 60 / -2`, -30.},
		{`return 60 * 2`, 120},
		{`return 60 * 2`, 120},
		{`return 60 * -2`, -120},
		{`return 60 * -2`, -120.0},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			res, err := exec(tt.in, c)
			r.NoError(err)
			r.Equal(fmt.Sprint(tt.exp), fmt.Sprint(res))

			f, err := ast.NewInteger(int(tt.exp))
			r.NoError(err)
			r.Equal(fmt.Sprint(tt.exp), f.String())
		})
	}
}

func Test_Integer_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `42`},
		{`%q`, `"42"`},
		{`%v`, `42`},
		{`%+v`, `42`},
		{`%#v`, `42`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.INT)

			r.Equal(tt.out, ft)
		})
	}
}
