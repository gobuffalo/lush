package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Var(t *testing.T) {
	table := []struct {
		in  string
		err bool
	}{
		{`a := "$"`, false},
		{`a := 42`, false},
		{`a := 3.14`, false},
		{`b := 3.14`, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			c.Set("b", 1)
			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
		})
	}
}
