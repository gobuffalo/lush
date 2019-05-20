package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
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
		{"%s", `foo := 42`},
		{"%q", `"foo := 42"`},
		{"%+v", assignv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			s := quick.VAR

			ft := fmt.Sprintf(tt.format, s)

			r.Equal(tt.out, ft)
		})
	}
}
