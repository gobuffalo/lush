package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Var(t *testing.T) {
	table := []struct {
		in  string
		err bool
	}{
		{`a := "$"`, false},
		{`a := 42`, false},
		{`a := 3.14`, false},
		{`b := 3.14`, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			c.Set("b", 1)
			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
		})
	}
}

func Test_Var_Format(t *testing.T) {
	assignv, err := jsonFixture("Var")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{"%s", `x := 1`},
		{"%q", `"x := 1"`},
		{"%+v", assignv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			id, err := ast.NewIdent([]byte("x"))
			r.NoError(err)

			v, err := ast.NewInteger(1)
			r.NoError(err)

			s, err := ast.NewVar(id, v)
			r.NoError(err)

			ft := fmt.Sprintf(tt.format, s)

			r.Equal(tt.out, ft)
		})
	}
}
