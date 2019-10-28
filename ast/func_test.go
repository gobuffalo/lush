package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
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

func Test_Func_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "func(foo) {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%q`, "\"func(foo) {\\n\\tfoo = 42\\n\\n\\tfoo := 42\\n}\""},
		{`%v`, "func(foo) {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%+v`, "func(foo) {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%#v`, "func(foo) {\n\tfoo = 42\n\n\tfoo := 42\n}"},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.FUNC)

			r.Equal(tt.out, ft)
		})
	}
}
