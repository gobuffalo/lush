package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Let(t *testing.T) {
	table := []struct {
		in  string
		exp interface{}
		err bool
	}{
		{`let a = "$"`, `$`, false},
		{"let a = `$`", `$`, false},
		{`let a = 42`, 42, false},
		{`let a = 3.14`, 3.14, false},
		{`let a = nil`, nil, false},
		{`let a = null`, nil, true},
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
