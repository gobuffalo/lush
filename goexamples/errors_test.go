package goexamples

import (
	"context"
	"testing"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_errorsExec(t *testing.T) {
	r := require.New(t)

	c := ast.NewContext(context.Background(), nil)

	s, err := lush.ParseFile("./errors.lush")
	r.NoError(err)
	r.True(Equal(c, s.Exec, errorsExec))
}
