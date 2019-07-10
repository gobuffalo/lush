package ast_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_Call_Helper(t *testing.T) {
	r := require.New(t)

	var res []interface{}
	c := NewContext()
	c.Set("foo", func(s interface{}) {
		res = append(res, s)
	})

	in := `
		foo("a")
		foo(42)
		foo(3.14)
	`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]interface{}{"a", 42, 3.14}, res)
}

func Test_Call_Helper_MultiArg(t *testing.T) {
	r := require.New(t)

	var res []string
	var ind []int
	c := NewContext()
	c.Set("foo", func(s string, i int) {
		res = append(res, s)
		ind = append(ind, i)
	})

	in := `
		foo("a", 0)
		foo("b", 1)
		foo("c", 2)
	`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]string{"a", "b", "c"}, res)
	r.Equal([]int{0, 1, 2}, ind)
}

func Test_Call_Helper_Let(t *testing.T) {
	r := require.New(t)

	in := `let x = func(i) {
		print(i)
	}
	x(42)
	`

	var ind []int
	c := NewContext()
	c.Set("print", func(i int) {
		ind = append(ind, i)
	})
	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]int{42}, ind)
}

func Test_Call_Helper_Let_MultiArgs(t *testing.T) {
	r := require.New(t)

	in := `let x = func(i, s) {
		print(i, s)
	}
	x(4.2, "foo")
	`

	var res []string
	var ind []float64
	c := NewContext()
	c.Set("print", func(i float64, s string) {
		ind = append(ind, i)
		res = append(res, s)
	})

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]float64{4.2}, ind)
	r.Equal([]string{"foo"}, res)
}

func Test_Call_Helper_Let_MultiArgs_Variadic(t *testing.T) {
	r := require.New(t)

	in := `let x = func(i, s) {
		print(i, s, s, s)
	}
	x(4.2, "foo")
	`

	var res []string
	var ind []float64
	c := NewContext()
	c.Set("print", func(i float64, s ...string) {
		ind = append(ind, i)
		res = append(res, s...)
	})

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]float64{4.2}, ind)
	r.Equal([]string{"foo", "foo", "foo"}, res)
}

func Test_Call_Helper_Let_MissingArgs(t *testing.T) {
	r := require.New(t)

	in := `let x = func(i, s) {
		print(i, s)
	}
	x(4.2)
	`

	var res []string
	var ind []float64
	c := NewContext()
	c.Set("print", func(i float64, s string) {
		ind = append(ind, i)
		res = append(res, s)
	})

	_, err := exec(in, c)
	r.Error(err)
}

type foo struct{}

func (f foo) Bar(i int, s string) string {
	return fmt.Sprintf("got %d/%s", i, s)
}

func (f foo) Vary(i int, s ...string) string {
	return fmt.Sprintf("got %d/%s", i, s)
}

func (f foo) WithFunc(fn func(i int, s string) string) string {
	return fn(42, "hello")
}

func Test_Call_External_Func(t *testing.T) {
	r := require.New(t)

	c := NewContext()
	c.Set("foo", foo{})

	var res string
	c.Set("capture", func(s string) {
		res = s
	})

	in := `capture(foo.Bar(1, "hi"))`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal(res, "got 1/hi")
}

func Test_Call_External_Func_Variadic(t *testing.T) {
	r := require.New(t)

	c := NewContext()
	c.Set("foo", foo{})

	var res string
	c.Set("capture", func(s string) {
		res = s
	})

	in := `capture(foo.Vary(1, "hi", "hello"))`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal(res, "got 1/[hi hello]")
}

func Test_Call_External_Func_WithFunc(t *testing.T) {
	t.Skip("TODO")
	r := require.New(t)

	c := NewContext()
	c.Set("foo", foo{})

	var res string
	c.Set("capture", func(s string) {
		res = s
	})

	in := `capture(foo.WithFunc(func(i, s) {return "asdf"}))`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal(res, "got 1/hi")
}

func Test_Call_Helper_Context(t *testing.T) {
	r := require.New(t)

	var res []interface{}
	c := NewContext()
	c.Set("goes", "nose")
	c.Set("foo", func(s interface{}, help *ast.Context) {
		res = append(res, s)
		res = append(res, help.Value("goes"))
	})

	in := `
		foo("a")
		foo(42)
		foo(3.14)
	`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]interface{}{"a", "nose", 42, "nose", 3.14, "nose"}, res)
}

func Test_Call_Helper_Context_Block(t *testing.T) {
	r := require.New(t)

	c := NewContext()
	c.Set("foo", func(s interface{}, help *ast.Context) (string, error) {
		if help.Block == nil {
			return "", errors.New("no block!")
		}
		res, err := help.Block.Exec(help)
		if err != nil {
			return "", err
		}
		return fmt.Sprint(res), nil
	})

	in := `return foo("a") {
	return "B"
}`

	res, err := exec(in, c)
	r.NoError(err)
	r.Equal("B", res.Value)
}

func Test_Call_Helper_Context_Block_With_Context(t *testing.T) {
	r := require.New(t)

	c := NewContext()
	c.Set("foo", func(s interface{}, help *ast.Context) (string, error) {
		if help.Block == nil {
			return "", errors.New("no block!")
		}
		nc := c.Clone()
		nc.Set("b", "BB")
		res, err := help.Block.Exec(nc)
		if err != nil {
			return "", err
		}
		return fmt.Sprint(res), nil
	})

	in := `
		return foo("a") {
			return b
		}
	`

	res, err := exec(in, c)
	r.NoError(err)
	r.Equal("BB", res.Value)
}

func Test_Call_Helper_Context_Block_Error(t *testing.T) {
	r := require.New(t)

	c := NewContext()
	c.Set("error", errors.New)
	c.Set("foo", func(s interface{}, help *ast.Context) (string, error) {
		if help.Block == nil {
			return "", errors.New("no block!")
		}
		res, err := help.Block.Exec(help)
		if err != nil {
			return "", err
		}
		return fmt.Sprint(res), nil
	})

	in := `
		return foo("a") {
			return error("oops")
		}
	`

	_, err := exec(in, c)
	r.Error(err)
}

func Test_Call_Helper_Options(t *testing.T) {
	r := require.New(t)

	var res []interface{}
	c := NewContext()
	c.Set("foo", func(s interface{}, opts map[string]interface{}) {
		res = append(res, s)
		if d, ok := opts["deep"]; ok {
			res = append(res, d)
		}
	})

	in := `
		foo("a", {deep: true})
		foo(42, {"deep": "prize"})
		foo(3.14)
	`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]interface{}{"a", true, 42, "prize", 3.14}, res)
}

func Test_Call_Helper_Options_Context(t *testing.T) {
	r := require.New(t)

	var res []interface{}
	c := NewContext()
	c.Set("goes", "nose")
	c.Set("foo", func(s interface{}, opts map[string]interface{}, help *ast.Context) {
		res = append(res, s)
		res = append(res, help.Value("goes"))
		if d, ok := opts["deep"]; ok {
			res = append(res, d)
		}
	})

	in := `
		foo("a", {deep: true})
		foo(42, {"deep": "prize"})
		foo(3.14)
	`

	_, err := exec(in, c)
	r.NoError(err)
	r.Equal([]interface{}{"a", "nose", true, 42, "nose", "prize", 3.14, "nose"}, res)
}

func Test_Call(t *testing.T) {
	brv, err := jsonFixture("Call")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "foo.Bar(42, 3.14, \"hi\") {\n\tfoo = 42\n\n\tfoo := 42\n}"},
		{`%q`, "\"foo.Bar(42, 3.14, \\\"hi\\\") {\\n\\tfoo = 42\\n\\n\\tfoo := 42\\n}\""},
		{`%+v`, brv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.CALL)

			r.Equal(tt.out, ft)
		})
	}
}

func Test_Call_Format(t *testing.T) {
	r := require.New(t)

	in := `create_table("users") {
	t.Column("id", "uuid", {"primary": true})

	t.Column("name", "string", {})

	t.Column("slug", "string", {})

	t.Column("email", "string", {})

	t.Column("token", "string", {"null": true})

	t.Timestamps()
}`

	p, err := parse(in)
	r.NoError(err)

	r.Equal(strings.TrimSpace(in), strings.TrimSpace(p.String()))
}
