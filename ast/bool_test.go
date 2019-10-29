package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Bool_Format(t *testing.T) {
	boolv, err := jsonFixture("Bool")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `true`},
		{`%v`, `true`},
		{`%#v`, `true`},
		{`%+v`, boolv},
		{`%q`, `"true"`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, ast.True)

			r.Equal(tt.out, ft)
		})
	}
}
