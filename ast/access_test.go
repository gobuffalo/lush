package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_Access_Array(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`a := [1,2,3]
		return a[1]`, 2, false},
		{`a := [1,2,3]
		return a[42]`, nil, true},
		{`return myArray[1]`, "petty", false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			c.Set("myArray", []string{"tom", "petty"})

			res, err := exec(tt.in, c)
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

func Test_Access_Map(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`a := {"foo": "bar", "baz": 2}
		return a["baz"]`, 2, false},
		{`a := {"foo": "bar", "baz": 2}
		return a["oops"]`, nil, true},
		{`return myMap["tom"]`, "petty", false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			c.Set("myMap", map[string]string{
				"tom":   "petty",
				"heart": "breakers",
			})

			res, err := exec(tt.in, c)
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

func Test_Access_Format(t *testing.T) {
	accessv, err := jsonFixture("Access")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, `foo[42]`},
		{`%q`, `"foo[42]"`},
		{`%+v`, accessv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)
			ft := fmt.Sprintf(tt.format, quick.ACCESS)

			r.Equal(tt.out, ft)
		})
	}
}
