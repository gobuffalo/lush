package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Return(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{in: `return 1`, out: 1},
		{in: `return "foo"`, out: "foo"},
		{in: `return true`, out: true},
		{in: `return [1, 2, 3]`, out: []interface{}{1, 2, 3}},
		{in: `return {"a": "b", "c": 3}`, out: map[interface{}]interface{}{
			"a": "b",
			"c": 3,
		}},
		{in: `return func() {return 1}()`, out: 1},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			res, err := exec(tt.in, NewContext())
			if tt.err {
				r.Error(err)
				return
			}

			r.NoError(err)
			r.NotNil(res)
			r.Equal(tt.out, res.Value)
		})
	}
}
