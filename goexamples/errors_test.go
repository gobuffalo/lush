package goexamples

import (
	"bytes"
	"context"
	"testing"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_errorsExec(t *testing.T) {
	r := require.New(t)

	bb := &bytes.Buffer{}

	c := ast.NewContext(context.Background(), bb)

	_, act := errorsExec(c)
	r.Error(act)

	s, err := lush.ParseFile("./errors.lush")
	r.NoError(err)

	_, exp := s.Exec(c)
	r.Error(exp)

	r.Equal(exp, act)
}
