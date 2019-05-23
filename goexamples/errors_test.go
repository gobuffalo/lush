package goexamples

import (
	"bytes"
	"context"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_errorsExec(t *testing.T) {
	r := require.New(t)

	bb := &bytes.Buffer{}

	c := ast.NewContext(context.Background(), bb)

	_, err := errorsExec(c)
	r.Error(err)
}
