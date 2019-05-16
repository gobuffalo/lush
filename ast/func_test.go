package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Func_Immediate_Call(t *testing.T) {
	r := require.New(t)

	in := `return func(x) {
		return x
	}("hi")`

	res, err := exec(in, NewContext())
	r.NoError(err)

	r.NotNil(res)
	r.Equal("hi", res.Value)
}
