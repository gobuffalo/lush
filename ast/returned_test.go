package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Returned_Return(t *testing.T) {
	table := []struct {
		in  interface{}
		out interface{}
		err bool
	}{
		{1, 1, false},
		{nil, nil, false},
		{abc, abc, false},
		{beatles, beatles, false},
		{fmt.Errorf("oops!"), nil, true},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt.in), func(st *testing.T) {
			r := require.New(st)

			ret := ast.NewReturned(tt.in)

			if tt.err {
				r.Error(ret.Err())
				return
			}

			r.NoError(ret.Err())
			r.Equal(tt.out, ret.Value)
		})
	}
}
