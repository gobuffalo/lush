package ast_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/stretchr/testify/require"
)

func Test_Comment(t *testing.T) {
	table := []struct {
		in  string
		out string
		err bool
	}{
		{`# foo`, `// foo`, false},
		{`#foo`, `// foo`, false},
		{`// foo`, `// foo`, false},
		{`//foo`, `// foo`, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			s, err := parse(tt.in)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(strings.TrimSpace(tt.out), strings.TrimSpace(s.String()))
		})
	}
}

func Test_Comment_MultiLine(t *testing.T) {
	r := require.New(t)

	in := `// for x := range [1, 2, 3] {
// 	fmt.Println(x)
//
// 	fmt.Println(x)
// }`

	c, err := ast.NewComment([]byte(in))
	r.NoError(err)
	r.Equal(in, c.String())
}

func Test_Comment_Format(t *testing.T) {
	ctv, err := jsonFixture("Comment")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `// hello`},
		{`%q`, `"// hello"`},
		{`%v`, `// hello`},
		{`%+v`, ctv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, ast.Comment{Value: "hello"})

			r.Equal(tt.out, ft)
		})
	}
}
