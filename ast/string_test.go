package ast_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_String(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{`let a = "hi!"`, `"hi!"`},
		{`let a = "hi! \"don't\" mind me!"`, `"hi!"`},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			_, err := exec(tt.in, c)
			r.NoError(err)

			x := c.Value("a")
			r.NotNil(x)
			r.Equal(tt.out, tt.out)
		})
	}
}

func Test_NewString(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{`"hi!"`, `"hi!"`},
		{"`hi!`", "`hi!`"},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			s, err := ast.NewString([]byte(tt.in))
			r.NoError(err)
			r.Equal(tt.out, s.String())
			un, err := strconv.Unquote(tt.out)
			r.NoError(err)
			r.Equal(un, s.Original)
		})
	}
}

func Test_String_Format(t *testing.T) {
	table := []struct {
		in     string
		format string
		out    string
	}{
		{`"hi"`, `%s`, `"hi"`},
		{`"hi"`, `%q`, "`\"hi\"`"},
		{`"hi"`, `%v`, `"hi"`},
		{`"hi"`, `%+v`, stringv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s_%s", tt.in, tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			s, err := ast.NewString([]byte(tt.in))
			r.NoError(err)

			ft := fmt.Sprintf(tt.format, s)

			r.Equal(tt.out, ft)
		})
	}
}

const stringv = `{
  "ast.String": {
    "Original": "hi",
    "QuoteFormat": "%q",
    "Meta": {
      "Filename": "",
      "Line": 0,
      "Col": 0,
      "Offset": 0,
      "Original": ""
    }
  }
}`
