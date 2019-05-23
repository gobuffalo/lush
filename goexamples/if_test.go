package goexamples

import (
	"bytes"
	"context"
	"testing"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_ifExec(t *testing.T) {
	r := require.New(t)

	bb1 := &bytes.Buffer{}

	c := ast.NewContext(context.Background(), bb1)

	ret, err := ifExec(c)
	r.NoError(err)
	r.Nil(ret)

	s, err := lush.ParseFile("./if.lush")
	r.NoError(err)

	bb2 := &bytes.Buffer{}

	c = ast.NewContext(context.Background(), bb2)

	ret, err = s.Exec(c)
	r.NoError(err)
	r.Nil(ret)

	r.Equal(bb2.String(), bb1.String())
}
