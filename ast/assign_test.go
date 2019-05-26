package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
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
		{`foo[0] = 42
		return foo`, []int{42}, false},
		{`baz["bar"] = 42
		return baz`, map[string]interface{}{"bar": 42}, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()

			c.Set("foo", []int{2})
			c.Set("baz", map[string]interface{}{})

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

func Test_Assign_Format(t *testing.T) {
	assignv, err := jsonFixture("Assign")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{"%s", `foo = 42`},
		{"%q", `"foo = 42"`},
		{"%+v", assignv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.ASSIGN)

			r.Equal(tt.out, ft)
		})
	}
}
