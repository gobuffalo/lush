package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Assign(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`x := 0
		x = 1
		return x`, 1, false},
		{`x = 1
		return x`, nil, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()

			res, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)
			r.Equal(tt.out, res.Value)
		})
	}
}

func Test_Assign_Up_Scope(t *testing.T) {
	r := require.New(t)

	in := `x := 0
func() {
	if true {
		x = 42
	}
}()
return x`

	c := NewContext()

	res, err := exec(in, c)
	r.NoError(err)
	r.Equal(42, res.Value)
}
