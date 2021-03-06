package ast_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_Map_String(t *testing.T) {
	r := require.New(t)

	in := `let j = {"a": "b", "foo": "bar", "h": 1, "y": func(x) {}}`

	s, err := parse(in)
	r.NoError(err)

	r.Equal(in, strings.TrimSpace(s.String()))
}

func Test_Map_Format(t *testing.T) {
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `{"foo": 42}`},
		{`%q`, `"{\"foo\": 42}"`},
		{`%v`, `{"foo": 42}`},
		{`%+v`, `{"foo": 42}`},
		{`%#v`, `map[string]interface {}{"foo":42}`},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.MAP)

			r.Equal(tt.out, ft)
		})
	}
}
