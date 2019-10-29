package ast_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_For_Array(t *testing.T) {
	table := []struct {
		in  string
		arr []interface{}
		out string
		err bool
	}{
		{`for (i) in a { capture(i) }`, abc, `abc`, false},
		{`for (i,x) in a { capture(i, x) }`, abc, `0a1b2c`, false},
		{`for i := range a { capture(i) }`, abc, `012`, false},
		{`for i, x := range a { capture(i, x) }`, abc, `0a1b2c`, false},
		{`for (i) in a { capture(i) break }`, abc, `a`, false},
		{`for (i,x) in a { capture(i, x) break }`, abc, `0a`, false},
		{`for i := range a { capture(i) break }`, abc, `0`, false},
		{`for i, x := range a { capture(i, x) break }`, abc, `0a`, false},
		{`for (i) in a { capture(i) continue capture("x") }`, abc, `abc`, false},
		{`for (i,x) in a { capture(i, x) continue capture("x") }`, abc, `0a1b2c`, false},
		{`for i := range a { capture(i) continue capture("x") }`, abc, `012`, false},
		{`for i, x := range a { capture(i, x) continue capture("x") }`, abc, `0a1b2c`, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			var res string
			c := NewContext()
			c.Set("a", tt.arr)
			c.Set("capture", func(i ...interface{}) {
				for _, x := range i {
					res += fmt.Sprint(x)
				}
			})

			_, err := exec(tt.in, c)
			r.NoError(err)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.out, res)
		})
	}
}

func Test_For_Map(t *testing.T) {

	keys := []string{"john", "paul", "george", "ringo"}
	values := []string{"guitar", "bass", "guitar", "drums"}
	long := []string{"john/guitar", "paul/bass", "george/guitar", "ringo/drums"}
	table := []struct {
		in  string
		arr map[string]string
		out []string
		err bool
	}{
		{`for (i) in a { capture(i) }`, beatles, values, false},
		{`for (i,x) in a { capture(i, x) }`, beatles, long, false},
		{`for i := range a { capture(i) }`, beatles, keys, false},
		{`for i, x := range a { capture(i, x) }`, beatles, long, false},
		{`for (i) in a { capture(i) continue capture("x") }`, beatles, values, false},
		{`for (i,x) in a { capture(i, x) continue capture("x") }`, beatles, long, false},
		{`for i := range a { capture(i) continue capture("x") }`, beatles, keys, false},
		{`for i, x := range a { capture(i, x) continue capture("x") }`, beatles, long, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			var res []string
			c := NewContext()
			c.Set("a", tt.arr)

			c.Set("capture", func(i ...interface{}) {
				var x []string
				for _, n := range i {
					x = append(x, fmt.Sprint(n))
				}
				res = append(res, strings.Join(x, "/"))
			})

			_, err := exec(tt.in, c)
			r.NoError(err)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)

			sort.Strings(tt.out)
			sort.Strings(res)
			r.Equal(tt.out, res)
		})
	}
}

func Test_For_Respects_Returns(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
	}{
		{`for x := range [1,2,3] { return 1 }; return 2`, 1},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			res, err := exec(tt.in, NewContext())
			r.NoError(err)
			r.Equal(tt.out, res.Value)
		})
	}
}

func Test_For_Infinite(t *testing.T) {
	r := require.New(t)

	in := `
for {
	if i == nil{
		let i = 0
	}

	i = i+1

	if i == 4 {
		return i
	}
}
`

	c := NewContext()

	res, err := exec(in, c)
	r.NoError(err)
	r.Equal(4, res.Value)
}

func Test_For_Infinite_String(t *testing.T) {
	r := require.New(t)

	in := `for {return foo}`
	out := `for {
	return foo
}`

	s, err := parse(in)
	r.NoError(err)
	r.Equal(out, strings.TrimSpace(s.String()))
}

func Test_For_Format(t *testing.T) {
	blv, err := jsonFixture("For")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "for (i, n) in [1, 2, 3] {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%v`, "for (i, n) in [1, 2, 3] {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%#v`, "for (i, n) in [1, 2, 3] {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%+v`, blv},
		{`%q`, "\"for (i, n) in [1, 2, 3] {\\n\\tfoo = 42\\n\\n\\tfoo := 42\\n}\""},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.FOR)

			r.Equal(tt.out, ft)
		})
	}
}
