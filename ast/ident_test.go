package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ident(t *testing.T) {
	table := []struct {
		in  string
		err bool
	}{
		{`let a = 1`, false},
		{`let a1 = 1`, false},
		{`let ab1 = 1`, false},
		{`let ab1a = 1`, false},
		{`let ab1c2 = 1`, false},
		{`let a_1 = 1`, false},
		{`let a_1_b = 1`, false},
		{`let a_B_1_C_2 = 1`, false},
		{`let null = 1`, false},
		{`let "a" = 1`, true},
		{`let 123 = 1`, true},
		{`let nil = 1`, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
		})
	}
}
