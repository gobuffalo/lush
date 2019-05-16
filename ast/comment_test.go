package ast_test

import (
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
