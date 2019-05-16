package ast_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Map_String(t *testing.T) {
	r := require.New(t)

	in := `let j = {"a": "b", "foo": "bar", "h": 1, y: func(x) {}}`

	s, err := parse(in)
	r.NoError(err)

	r.Equal(in, strings.TrimSpace(s.String()))
}
