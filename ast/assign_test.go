package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Assign(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`x := 0
		x = 1
		return x`, 1, false},
		{`x = 1
		return x`, nil, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()

			res, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)
			r.Equal(tt.out, res.Value)
		})
	}
}

func Test_Assign_Up_Scope(t *testing.T) {
	r := require.New(t)

	in := `x := 0
func() {
	if true {
		x = 42
	}
}()
return x`

	c := NewContext()

	res, err := exec(in, c)
	r.NoError(err)
	r.Equal(42, res.Value)
}

func Test_Assign_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{"%s", `x = 1`},
		{"%q", `"x = 1"`},
		{"%+v", assignv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			id, err := ast.NewIdent([]byte("x"))
			r.NoError(err)

			v, err := ast.NewInteger(1)
			r.NoError(err)

			s, err := ast.NewAssign(id, v)
			r.NoError(err)

			ft := fmt.Sprintf(tt.format, s)

			r.Equal(tt.out, ft)
		})
	}
}

const assignv = `{
  "ast.Assign": {
    "Meta": {
      "Filename": "",
      "Line": 0,
      "Col": 0,
      "Offset": 0,
      "Original": ""
    },
    "Name": {
      "ast.Ident": {
        "Meta": {
          "Filename": "",
          "Line": 0,
          "Col": 0,
          "Offset": 0,
          "Original": ""
        },
        "Name": {
          "string": "x"
        }
      }
    },
    "Value": 1
  }
}`
