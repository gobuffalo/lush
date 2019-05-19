package ast_test

import (
	"fmt"
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
		in  []interface{}
		out string
	}{
		{[]interface{}{1, 2, 3}, "[1, 2, 3]"},
		{[]interface{}{newString("a"), newString("b")}, `["a", "b"]`},
		{[]interface{}{newString("a"), ast.Float(3.14), ast.Bool(true)}, `["a", 3.14, true]`},
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
	table := []struct {
		in     []interface{}
		format string
		out    string
	}{
		{[]interface{}{1, 2, 3}, "%s", `[1, 2, 3]`},
		{[]interface{}{1, 2, 3}, "%q", `"[1, 2, 3]"`},
		{[]interface{}{1, 2, 3}, "%+v", arrayv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s_%s", tt.in, tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			s, err := ast.NewArray(tt.in)
			r.NoError(err)

			ft := fmt.Sprintf(tt.format, s)

			r.Equal(tt.out, ft)
		})
	}
}

const arrayv = `{
  "ast.Array": {
    "Meta": {
      "Filename": "",
      "Line": 0,
      "Col": 0,
      "Offset": 0,
      "Original": ""
    },
    "Value": {
      "[]interface {}": [
        1,
        2,
        3
      ]
    }
  }
}`
