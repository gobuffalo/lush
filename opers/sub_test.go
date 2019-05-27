package opers

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/types"
	"github.com/stretchr/testify/require"
)

func Test_Sub(t *testing.T) {
	table := []struct {
		a   interface{}
		b   interface{}
		exp interface{}
		err bool
	}{
		{42, 2, 40, false},
		{map[string]interface{}{"a": "A", "b": "B"}, "b", map[string]interface{}{"a": "A"}, false},
		{types.Mapper(map[string]interface{}{"a": "A", "b": "B"}), "b", map[string]interface{}{"a": "A"}, false},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt), func(st *testing.T) {
			r := require.New(st)

			res, err := Sub(tt.a, tt.b)

			if tt.err {
				r.Error(err)
			}
			r.NoError(err)
			r.Equal(tt.exp, res)
		})
	}
}
