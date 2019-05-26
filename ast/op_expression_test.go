package ast_test

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/ast/internal/quick"
	"github.com/stretchr/testify/require"
)

func Test_OpExpression_Equal(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if ("a" == "a") {return true} return false`, true, false},
		{`if (42 == 42) {return true} return false`, true, false},
		{`if (42 == magic) {return true} return false`, true, false},
		{`if (3.14 == 3.14) {return true} return false`, true, false},
		{`if (true == true) {return true} return false`, true, false},
		{`if ([1,2,3] == [1,2,3]) {return true} return false`, true, false},
		{`if ({a: "A"} == {a: "A"}) {return true} return false`, true, false},
		{`if ("a" == "b") {return true} return false`, false, false},
		{`if (42 == 4.2) {return true} return false`, false, false},
		{`if (3.14 == 314) {return true} return false`, false, false},
		{`if (true == false) {return true} return false`, false, false},
		{`if ([1,2,3] == [3,2,1]) {return true} return false`, false, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			c.Set("magic", ast.Integer{Value: 42})
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

func Test_OpExpression_NotEqual(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if ("a" != "a") {return true} return false`, false, false},
		{`if (42 != 42) {return true} return false`, false, false},
		{`if (3.14 != 3.14) {return true} return false`, false, false},
		{`if (true != true) {return true} return false`, false, false},
		{`if ([1,2,3] != [1,2,3]) {return true} return false`, false, false},
		{`if ({a: "A"} != {a: "A"}) {return true} return false`, false, false},
		{`if ("a" != "b") {return true} return false`, true, false},
		{`if (42 != 4.2) {return true} return false`, true, false},
		{`if (3.14 != 314) {return true} return false`, true, false},
		{`if (true != false) {return true} return false`, true, false},
		{`if ([1,2,3] != [3,2,1]) {return true} return false`, true, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_Regexp(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if ("a" ~= "a") {return true} return false`, true, false},
		{`if (42 ~= 42) {return true} return false`, true, false},
		{`if (4.2 ~= 42) {return true} return false`, false, false},
		{`if (3.14 ~= 3.14) {return true} return false`, true, false},
		{`if (true ~= true) {return true} return false`, true, false},
		{`if (true ~= trUe) {return true} return false`, false, false},
		{`if ([1,2,3] ~= [1,2,3]) {return true} return false`, false, true},
		{`if ({a: "A"} ~= {a: "A"}) {return true} return false`, false, true},
		{`if ("a" ~= "b") {return true} return false`, false, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_Add(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`return 4 + 2`, 6, false},
		{`return 4 + 3.50`, 7.50, false},
		{`return 4 + -3.50`, 0.50, false},
		{`return 3.50 + 4`, 7.50, false},
		{`return "a" + "b"`, "ab", false},
		{`return 10 + ( 4 + 2 )`, 16, false},
		{`return ( ( "a" + "b" ) + ("c" + "d") )`, "abcd", false},
		{`return [1] + [2]`, []interface{}{1, 2}, false},
		{`return y + " " + z`, "zoo farm", false},
		{`return 4 + "a"`, nil, true},
		{`return "a" + 4`, nil, true},
		{`return {a: 1, b: "hi"} + {c: true}`, nil, true},
		{`return "a" + "b" + "c"`, "abc", false},
		{`return ( "a" + "b" ) + ("c" + "d")`, "abcd", true}, // TODO
		{`return ( ( "a" + "b" ) + "c" + "d" )`, nil, true},  // TODO
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			c.Set("y", "zoo")
			c.Set("z", "farm")

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

func Test_OpExpression_Sub(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`return 4 - 2`, 2, false},
		{`return 4 - 3.50`, 0.50, false},
		{`return 4 - -3.50`, 7.50, false},
		{`return 3.50 - 4`, -0.50, false},
		{`return 10 - ( 4 - 2 )`, 8, false},
		{`return "a" - "b"`, nil, true},
		{`return ( ( "a" - "b" ) - ("c" - "d") )`, nil, true},
		{`return [1] - [2]`, nil, true},
		{`return 4 - "a"`, nil, true},
		{`return "a" - 4`, nil, true},
		{`return {a: 1, b: "hi"} - {c: true}`, nil, true},
		{`return "a" - "b" - "c"`, "abc", true},
		{`return ( "a" - "b" ) - ("c" - "d")`, "abcd", true},
		{`return ( ( "a" - "b" ) - "c" - "d" )`, nil, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_Multiply(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`return 4 * 2`, 8, false},
		{`return 4 * 3.50`, 14.0, false},
		{`return 4 * -3.50`, -14.0, false},
		{`return 3.50 * 4`, 14.0, false},
		{`return 10 * ( 4 * 2 )`, 80, false},
		{`return "a" * "b"`, nil, true},
		{`return ( ( "a" * "b" ) * ("c" * "d") )`, nil, true},
		{`return [1] * [2]`, nil, true},
		{`return 4 * "a"`, nil, true},
		{`return "a" * 4`, nil, true},
		{`return {a: 1, b: "hi"} * {c: true}`, nil, true},
		{`return "a" * "b" * "c"`, "abc", true},
		{`return ( "a" * "b" ) * ("c" * "d")`, "abcd", true},
		{`return ( ( "a" * "b" ) * "c" * "d" )`, nil, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_Divide(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`return 4 / 2`, 2, false},
		{`return 4 / 2.0`, 2.0, false},
		{`return 4 / -2.0`, -2.0, false},
		{`return 3.50 / 4`, 0.875, false},
		{`return 10 / ( 4 / 2 )`, 5, false},
		{`return "a" / "b"`, nil, true},
		{`return ( ( "a" / "b" ) / ("c" / "d") )`, nil, true},
		{`return [1] / [2]`, nil, true},
		{`return 4 / "a"`, nil, true},
		{`return "a" / 4`, nil, true},
		{`return {a: 1, b: "hi"} / {c: true}`, nil, true},
		{`return "a" / "b" / "c"`, "abc", true},
		{`return ( "a" / "b" ) / ("c" / "d")`, "abcd", true},
		{`return ( ( "a" / "b" ) / "c" / "d" )`, nil, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_Modulus(t *testing.T) {
	table := []struct {
		in  string
		out interface{}
		err bool
	}{
		{`return 4 % 3`, 1, false},
		{`return 4 % 2.0`, nil, true},
		{`return 4 % -2.0`, nil, true},
		{`return 3.50 % 4`, nil, true},
		{`return 10 % ( 4 % 2 )`, 0, false},
		{`return "a" % "b"`, nil, true},
		{`return ( ( "a" % "b" ) % ("c" % "d") )`, nil, true},
		{`return [1] % [2]`, nil, true},
		{`return 4 % "a"`, nil, true},
		{`return "a" % 4`, nil, true},
		{`return {a: 1, b: "hi"} % {c: true}`, nil, true},
		{`return "a" % "b" % "c"`, "abc", true},
		{`return ( "a" % "b" ) % ("c" % "d")`, "abcd", true},
		{`return ( ( "a" % "b" ) % "c" % "d" )`, nil, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_LessThan(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if ("a" < "b") {return true} return false`, true, false},
		{`if (42 < 43) {return true} return false`, true, false},
		{`if (3.14 < 3.15) {return true} return false`, true, false},
		{`if ("a" < "a") {return true} return false`, false, false},
		{`if (42 < 42) {return true} return false`, false, false},
		{`if (3.14 < 3.14) {return true} return false`, false, false},

		{`if (true < true) {return true} return false`, false, true},
		{`if ([1,2,3] < [3,2,1]) {return true} return false`, false, true},
		{`if ({a: "A"} < {a: "A"}) {return true} return false`, false, true},
		{`if (true < false) {return true} return false`, false, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_GreaterThan(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		// {`if ("b" > "a") {return true} return false`, true, false},
		// {`if (43 > 42) {return true} return false`, true, false},
		// {`if (3.15 > 3.14) {return true} return false`, true, false},
		{`if ("a" > "a") {return true} return false`, false, false},
		// {`if (42 > 42) {return true} return false`, false, false},
		// {`if (3.14 > 3.14) {return true} return false`, false, false},
		//
		// {`if (true > true) {return true} return false`, false, true},
		// {`if ([3,2,1] > [1,2,3]) {return true} return false`, true, true},
		// {`if ({a: "A"} > {a: "A"}) {return true} return false`, false, true},
		// {`if (true > false) {return true} return false`, false, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_LessThanEqualTo(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if ("a" <= "b") {return true} return false`, true, false},
		{`if ("a" <= "a") {return true} return false`, true, false},
		{`if (42 <= 43) {return true} return false`, true, false},
		{`if (42 <= 42) {return true} return false`, true, false},
		{`if (3.14 <= 3.15) {return true} return false`, true, false},
		{`if (3.14 <= 3.14) {return true} return false`, true, false},
		{`if ("b" <= "a") {return true} return false`, false, false},
		{`if (43 <= 42) {return true} return false`, false, false},
		{`if (3.15 <= 3.14) {return true} return false`, false, false},

		{`if (true <= true) {return true} return false`, true, true},
		{`if (false <= true) {return true} return false`, true, true},
		{`if ([1,2,3] <= [3,2,1]) {return true} return false`, true, true},
		{`if ([1,2,3] <= [1,2,3]) {return true} return false`, true, true},
		{`if ({a: "A"} <= {a: "A"}) {return true} return false`, true, true},
		{`if (true <= false) {return true} return false`, false, true},
		{`if ([1,2,4] <= [1,2,3]) {return true} return false`, false, true},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_GreaterThanEqualTo(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if ("a" >= "a") {return true} return false`, true, false},
		{`if (42 >= 42) {return true} return false`, true, false},
		{`if (43 >= 42) {return true} return false`, true, false},
		{`if ("b" >= "a") {return true} return false`, true, false},
		{`if (3.15 >= 3.14) {return true} return false`, true, false},
		{`if (true >= false) {return true} return false`, true, false},
		{`if ([1,2,4] >= [1,2,3]) {return true} return false`, true, false},
		{`if ({a: "A"} >= {a: "A"}) {return true} return false`, true, false},
		{`if (3.14 >= 3.14) {return true} return false`, true, false},
		{`if (3.15 >= 3.14) {return true} return false`, true, false},
		{`if (true >= true) {return true} return false`, true, false},
		{`if ("a" >= "b") {return true} return false`, false, false},
		{`if (42 >= 43) {return true} return false`, false, false},
		{`if (3.14 >= 3.15) {return true} return false`, false, false},
		{`if (false >= true) {return true} return false`, false, false},
		{`if ([1,2,3] >= [3,2,1]) {return true} return false`, false, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()

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

func Test_OpExpression_And(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if (true && true) {return true} return false`, true, false},
		{`if (n && true) {return true} return false`, true, false},
		{`if ((1 == 1) && true) {return true} return false`, true, false},
		{`if ((1 == 1) && (2 == 2)) {return true} return false`, true, false},
		{`if ("a" && "b") {return true} return false`, true, false},
		{`if (tom && "b") {return true} return false`, true, false},
		{`if ([1] && [2]) {return true} return false`, true, false},
		{`if ({"a": "A"} && {"b": "B"}) {return true} return false`, true, false},
		{`if (filledArray && "b") {return true} return false`, true, false},
		{`if (filledMap && "b") {return true} return false`, true, false},
		{`if (true && false) {return true} return false`, false, false},
		{`if (nil && false) {return true} return false`, false, false},
		{`if ("" && "") {return true} return false`, false, false},
		{`if ((1 == 0) && (2 == 2)) {return true} return false`, false, false},
		{`if ([] && []) {return true} return false`, false, false},
		{`if ({} && {}) {return true} return false`, false, false},
		{`if (1 && nil) {return true} return false`, false, false},
		{`if (nil && nil) {return true} return false`, false, false},
		{`if (empty && "b") {return true} return false`, false, false},
		{`if (emptyArray && "b") {return true} return false`, false, false},
		{`if (emptyMap && "b") {return true} return false`, false, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			c.Set("tom", "petty")
			c.Set("empty", "")
			c.Set("n", 0)
			c.Set("emptyArray", []int{})
			c.Set("emptyMap", map[int]int{})
			c.Set("filledArray", []int{1})
			c.Set("filledMap", map[int]int{1: 2})
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

func Test_OpExpression_Or(t *testing.T) {
	table := []struct {
		in  string
		out bool
		err bool
	}{
		{`if (true || true) {return true} return false`, true, false},
		{`if (n || true) {return true} return false`, true, false},
		{`if ((1 == 1) || true) {return true} return false`, true, false},
		{`if ((1 == 1) || (2 == 2)) {return true} return false`, true, false},
		{`if ("a" || "b") {return true} return false`, true, false},
		{`if (tom || "b") {return true} return false`, true, false},
		{`if ([1] || [2]) {return true} return false`, true, false},
		{`if ({"a": "A"} || {"b": "B"}) {return true} return false`, true, false},
		{`if (filledArray || "b") {return true} return false`, true, false},
		{`if (filledMap || "b") {return true} return false`, true, false},
		{`if (true || false) {return true} return false`, true, false},
		{`if (nil || false) {return true} return false`, false, false},
		{`if ("" || "") {return true} return false`, false, false},
		{`if ((1 == 0) || (2 == 2)) {return true} return false`, true, false},
		{`if ([] || []) {return true} return false`, false, false},
		{`if ({} || {}) {return true} return false`, false, false},
		{`if (1 || nil) {return true} return false`, true, false},
		{`if (nil || nil) {return true} return false`, false, false},
		{`if (empty || "b") {return true} return false`, true, false},
		{`if (emptyArray || "b") {return true} return false`, true, false},
		{`if (emptyMap || "b") {return true} return false`, true, false},
	}

	for _, tt := range table {
		t.Run(tt.in, func(st *testing.T) {
			r := require.New(st)

			c := NewContext()
			c.Set("tom", "petty")
			c.Set("empty", "")
			c.Set("n", 0)
			c.Set("emptyArray", []int{})
			c.Set("emptyMap", map[int]int{})
			c.Set("filledArray", []int{1})
			c.Set("filledMap", map[int]int{1: 2})
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

func Test_OpExpression_OpExpressionmat(t *testing.T) {
	blv, err := jsonFixture("OpExpression")
	if err != nil {
		t.Fatal(err)
	}
	table := []struct {
		format string
		out    string
	}{
		{`%s`, "(42 == 3.14)"},
		{`%q`, "\"(42 == 3.14)\""},
		{`%+v`, blv},
	}

	for _, tt := range table {
		t.Run(fmt.Sprintf("%s_%s", tt.format, tt.out), func(st *testing.T) {
			r := require.New(st)

			ft := fmt.Sprintf(tt.format, quick.OPEXPR)

			r.Equal(tt.out, ft)
		})
	}
}
