package ast_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/gobuffalo/lush/builtins"
	"github.com/stretchr/testify/require"
)

func Test_Import(t *testing.T) {
	table := []struct {
		in  string
		out string
		err bool
	}{
		{`import "fmt"; fmt.Print("hi")`, "hi", false},
		{`fmt.Print("hi")`, "", true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			bb := &bytes.Buffer{}

			c := NewContext()
			c.Imports.Store("fmt", builtins.NewFmt(bb))

			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)
			r.Equal(tt.out, bb.String())
		})
	}
}

func Test_Import_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `import "fmt"`},
		{`%q`, `"import \"fmt\""`},
		{`%v`, `import "fmt"`},
		{`%+v`, `import "fmt"`},
		{`%#v`, `import "fmt"`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.IMPORT)

			r.Equal(tt.out, ft)
		})
	}
}
