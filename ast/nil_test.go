package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Nil(t *testing.T) {
	r := require.New(t)

	in := `return nil`

	c := NewContext()
	res, err := exec(in, c)
	r.NoError(err)
	r.Equal(nil, res.Value)
}

func Test_Nil_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `nil`},
		{`%v`, `nil`},
		{`%+v`, `nil`},
		{`%#v`, `nil`},
		{`%q`, `"nil"`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, ast.Nil{})

			r.Equal(tt.out, ft)
		})
	}
}
