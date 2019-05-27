package add

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/types"
	"github.com/stretchr/testify/require"
)

func Test_Int(t *testing.T) {
	table := []struct {
		a   int
		b   interface{}
		out interface{}
		err bool
	}{
		{1, 2, 3, false},
		{1, 2.0, 3.0, false},
		{1, types.Floater(2.0), 3.0, false},
		{1, types.Integer(2), 3, false},
		{1, "nope", 0, true},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt), func(st *testing.T) {
			r := require.New(st)

			res, err := Int(tt.a, tt.b)

			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)

			r.Equal(tt.out, res)

		})
	}
}
