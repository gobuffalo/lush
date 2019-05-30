package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Current(t *testing.T) {
	r := require.New(t)

	in := `return current`

	exp := NewContext()

	ret, err := exec(in, exp)
	r.NoError(err)
	r.Equal(ret.Value, exp)
}
