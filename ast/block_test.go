package ast_test

import (
	"strings"
	"testing"

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
