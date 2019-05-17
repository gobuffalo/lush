package ast_test

import (
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Array_Interface(t *testing.T) {
	r := require.New(t)
	a := ast.Array{Value: abc}
	r.Equal(abc, a.Interface())
}

func Test_Array(t *testing.T) {
	table := []struct {
		in  string
		exp []interface{}
		err bool
	}{
		{`let a = [1,2,3]`, []interface{}{1, 2, 3}, false},
		{`let a = [1,["a", "b"],3]`, []interface{}{1, []interface{}{"a", "b"}, 3}, false},
		{`let a = [`, nil, true},
		{`let a = ["a", "b", 42]`, []interface{}{"a", "b", 42}, false},
		{`let a = [true, {"foo": "bar"}, 42]`, []interface{}{true, map[interface{}]interface{}{"foo": "bar"}, 42}, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)

			x := c.Value("a")
			r.NotNil(x)
			a, ok := x.([]interface{})
			r.True(ok)
			r.Equal(tt.exp, a)
		})
	}
}

func Test_Array_String(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{"return [1,2,3]", "[1, 2, 3]"},
		{`return ["a","b"]`, `["a", "b"]`},
		{`return [  "a"   ,   3.14,true   ]`, `["a", 3.14, true]`},
	}

	for _, tt := range table {
		t.Run(tt.out, func(st *testing.T) {
			r := require.New(st)

			p, err := parse(tt.in)
			r.NoError(err)
			r.Len(p.Statements, 1)

			ret, ok := p.Statements[0].(ast.Return)
			r.True(ok)

			r.Len(ret.Statements, 1)

			a, ok := ret.Statements[0].(ast.Array)
			r.True(ok)
			r.Equal(tt.out, a.String())

		})
	}
}
