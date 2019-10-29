package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/ast/internal/quick"
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
		{`let a = [true, {"foo": "bar"}, 42]`, []interface{}{true, map[string]interface{}{"foo": "bar"}, 42}, false},
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
	s1, _ := ast.NewString([]byte("a"))
	s2, _ := ast.NewString([]byte("b"))
	table := []struct {
		in  []interface{}
		out string
	}{
		{[]interface{}{1, 2, 3}, "[1, 2, 3]"},
		{[]interface{}{s1, s2}, `["a", "b"]`},
		{[]interface{}{s1, quick.FLOAT, ast.True}, `["a", 3.14, true]`},
	}

	for _, tt := range table {
		t.Run(tt.out, func(st *testing.T) {
			r := require.New(st)
			a, err := ast.NewArray(tt.in)
			r.NoError(err)
			r.Equal(tt.out, a.String())
		})
	}
}

func Test_Array_Format(t *testing.T) {
	arrayv, err := jsonFixture("Array")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{"%s", `[1, 2, 3]`},
		{"%v", `[1, 2, 3]`},
		{"%#v", `[1, 2, 3]`},
		{"%+v", arrayv},
		{"%q", `"[1, 2, 3]"`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.ARRAY)

			r.Equal(tt.out, ft)
		})
	}
}
