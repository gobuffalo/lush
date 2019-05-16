package ast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_If(t *testing.T) {
	table := []struct {
		in  string
		out string
		err bool
	}{
		{`if ("a" == "a") { capture("$") }`, "$", false},
		{`if ("a" == "b") { capture("$") }`, "", false},
		{`if (1 == 1) { capture("$") }`, "$", false},
		{`if (1 == 2) { capture("$") }`, "", false},
		{`if ([1,2,3] == [1,2,3]) { capture("$") }`, "$", false},
		{`if ([3,2, 1] == [1,2,3]) { capture("$") }`, "", false},
		{`if ("a" != "a") { capture("$") }`, "", false},
		{`if ("a" != "b") { capture("$") }`, "$", false},
		{`if (1 != 1) { capture("$") }`, "", false},
		{`if (1 != 2) { capture("$") }`, "$", false},
		{`if ([1,2,3] != [1,2,3]) { capture("$") }`, "", false},
		{`if ([3,2, 1] != [1,2,3]) { capture("$") }`, "$", false},
		{`if "a" == "a" { capture("$") }`, "$", false},
		{`if "a" == "b" { capture("$") }`, "", false},
		{`if 1 == 1 { capture("$") }`, "$", false},
		{`if 1 == 2 { capture("$") }`, "", false},
		{`if [1,2,3] == [1,2,3] { capture("$") }`, "$", false},
		{`if [3,2, 1] == [1,2,3] { capture("$") }`, "", false},
		{`if "a" != "a" { capture("$") }`, "", false},
		{`if "a" != "b" { capture("$") }`, "$", false},
		{`if 1 != 1 { capture("$") }`, "", false},
		{`if 1 != 2 { capture("$") }`, "$", false},
		{`if [1,2,3] != [1,2,3] { capture("$") }`, "", false},
		{`if [3,2, 1] != [1,2,3] { capture("$") }`, "$", false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)
			c := NewContext()
			var res string
			c.Set("capture", func(b string) {
				res = b
			})
			_, err := exec(tt.in, c)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.out, res)
		})
	}
}

func Test_If_Respects_Returns(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
	}{
		{`if true { return 1 }; return 2`, 1},
		{`if false { return 1 }; return 2`, 2},
		{`if false { return 1 } else { return 3}; return 2`, 3},
		{`if false { return 1 } else if true { return 4 } else { return 3}; return 2`, 4},
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

func Test_If_PreCondition(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
	}{
		{`if  x := 1 ; x == 1 { return 1 }; return 2`, 1},
		{`if let x = 1; x == 1 { return 1 }; return 2`, 1},
		{`if  x := 1 ; x == 2 { return 1 }; return 2`, 2},
		{`if let x = 1; x == 2 { return 1 }; return 2`, 2},
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
