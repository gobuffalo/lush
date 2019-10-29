package ast_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_Range_Array(t *testing.T) {
	table := []struct {
		in  string
		exp interface{}
		err bool
	}{
		{`for i := range a { capture(i) }`, []interface{}{0, 1, 2}, false},
		{`for i := range [1,2,3] { capture(i) }`, []interface{}{0, 1, 2}, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			c.Set("a", tt.exp)
			var res []interface{}
			c.Set("capture", func(i interface{}) {
				res = append(res, i)
			})

			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.exp, res)
		})
	}
}

func Test_Range_Array_MultiArg(t *testing.T) {
	table := []struct {
		in  string
		exp interface{}
		err bool
	}{
		{`for i, x := range a { capture(i, x) }`, []string{"0/1", "1/2", "2/3"}, false},
		{`for i,x:=range[1,2,3]{capture(i,x)}`, []string{"0/1", "1/2", "2/3"}, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			c.Set("a", []int{1, 2, 3})
			var res []string
			c.Set("capture", func(i interface{}, x interface{}) {
				res = append(res, fmt.Sprintf("%v/%v", i, x))
			})

			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.exp, res)
		})
	}
}

func Test_Range_Array_Break(t *testing.T) {
	r := require.New(t)

	in := `for i, n := range a {
		print(n)
		break
	}`

	c := NewContext()
	a := []int{4, 5, 6}
	c.Set("a", a)

	var res []int
	c.Set("print", func(i int) {
		res = append(res, i)
	})

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]int{4}, res)
}

func Test_Range_Array_Continue(t *testing.T) {
	r := require.New(t)

	in := `for i, n := range a {
		print(n)
		continue
		print(42)
	}`

	x, err := parse(in)
	r.NoError(err)

	c := NewContext()
	a := []int{4, 5, 6}
	c.Set("a", a)

	var res []int
	c.Set("print", func(i int) {
		res = append(res, i)
	})

	_, err = x.Exec(c)
	r.NoError(err)
	r.Equal(a, res)
}

func Test_Range_Array_String(t *testing.T) {
	r := require.New(t)

	in := `for i := range a {
	print(i)

	continue

	print(42)
}`

	n, err := parse(in)
	r.NoError(err)

	r.Equal(in, strings.TrimSpace(n.String()))
}

func Test_Range_Format(t *testing.T) {
	blv, err := jsonFixture("Range")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "for i, n := range [1, 2, 3] {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%v`, "for i, n := range [1, 2, 3] {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%#v`, "for i, n := range [1, 2, 3] {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%+v`, blv},
		{`%q`, "\"for i, n := range [1, 2, 3] {\\n\\tfoo = 42\\n\\n\\tfoo := 42\\n}\""},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.RANGE)

			r.Equal(tt.out, ft)
		})
	}
}
