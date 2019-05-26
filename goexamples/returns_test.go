package goexamples

import (
	"context"
	"testing"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_returnsExec(t *testing.T) {
	r := require.New(t)

	c := ast.NewContext(context.Background(), nil)

	s, err := lush.ParseFile("returns.lush")
	r.NoError(err)
	r.True(Equal(c, s.Exec, returnsExec))
}