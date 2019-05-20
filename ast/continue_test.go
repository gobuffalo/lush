package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Continue_Format(t *testing.T) {
	brv, err := jsonFixture("Continue")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `continue`},
		{`%q`, `"continue"`},
		{`%+v`, brv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, ast.Continue{})

			r.Equal(tt.out, ft)
		})
	}
}
