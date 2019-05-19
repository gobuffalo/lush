package ast_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Block_String(t *testing.T) {
	r := require.New(t)

	in := `
a := func(){
print(foo)
					if (    true        ) {return false}
}
`

	out := `
a := func() {
	print(foo)

	if true {
		return false
	}
}`
	s, err := parse(in)
	r.NoError(err)
	r.Equal(strings.TrimSpace(out), strings.TrimSpace(s.String()))
}

func Test_Block_format(t *testing.T) {
	blv, err := jsonFixture("Block")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "{\n\tx = 1\n}"},
		{`%q`, "\"{\\n\\tx = 1\\n}\""},
		{`%v`, "{\n\tx = 1\n}"},
		{`%+v`, blv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			id, err := ast.NewIdent([]byte("x"))
			r.NoError(err)

			v, err := ast.NewInteger(1)
			r.NoError(err)

			an, err := ast.NewAssign(id, v)
			r.NoError(err)

			s, err := ast.NewBlock(ast.Statements{an})
			r.NoError(err)

			ft := fmt.Sprintf(tt.format, s)

			r.Equal(tt.out, ft)
		})
	}
}
