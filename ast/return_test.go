package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
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
		{in: `return {"a": "b", "c": 3}`, out: map[string]interface{}{
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

func Test_Return_Format(t *testing.T) {
	blv, err := jsonFixture("Return")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "return 42"},
		{`%v`, "return 42"},
		{`%#v`, "return 42"},
		{`%+v`, blv},
		{`%q`, "\"return 42\""},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.RETURN)

			r.Equal(tt.out, ft)
		})
	}
}
