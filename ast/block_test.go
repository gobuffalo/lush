package ast_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
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
		{`%s`, "{\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%q`, "\"{\\n\\tfoo = 42\\n\\n\\tfoo := 42\\n}\""},
		{`%+v`, blv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.BLOCK)

			r.Equal(tt.out, ft)
		})
	}
}
