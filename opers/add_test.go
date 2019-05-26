package opers

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/types"
	"github.com/stretchr/testify/require"
)

type adder string

func (a adder) Add(b interface{}) (interface{}, error) {
	return fmt.Sprintf("%v/%v", a, b), nil
}

func Test_Add(t *testing.T) {
	table := []struct {
		A   interface{}
		B   interface{}
		Exp interface{}
		err bool
	}{
		{1, 1, 2, false},
		{1, 1.0, 2.0, false},
		{1, 1.0, 2.0, false},
		{1.0, 1.0, 2.0, false},
		{1.0, 1, 2.0, false},
		{"a", "b", "ab", false},
		{[]interface{}{"a"}, []interface{}{"b"}, []interface{}{"a", "b"}, false},
		{types.Stringer("a"), types.Stringer("b"), "ab", false},
		{types.Stringer("a"), "b", "ab", false},
		{"a", types.Stringer("b"), "ab", false},
		{adder("a"), adder("b"), "a/b", false},
		{types.Integer(1), types.Integer(1), 2, false},
		{types.Integer(1), 1, 2, false},
		{1, types.Integer(1), 2, false},
		{1, types.Floater(1.0), 2.0, false},
		{types.Floater(1.0), types.Floater(1.0), 2.0, false},
		{types.Floater(1.0), types.Integer(1), 2.0, false},
		{types.Floater(1.0), 1.0, 2.0, false},
		{1.0, types.Floater(1.0), 2.0, false},
		{types.Slicer([]interface{}{"a"}), types.Slicer([]interface{}{"b"}), []interface{}{"a", "b"}, false},
		{map[string]interface{}{"a": "A"}, map[string]interface{}{"b": "B"}, map[string]interface{}{"a": "A", "b": "B"}, false},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt), func(st *testing.T) {
			r := require.New(st)

			res, err := Add(tt.A, tt.B)
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)
			r.Equal(tt.Exp, res)
		})
	}
}
