package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Nil(t *testing.T) {
	r := require.New(t)

	in := `return nil`

	c := NewContext()
	res, err := exec(in, c)
	r.NoError(err)
	r.Equal(nil, res.Value)
}
