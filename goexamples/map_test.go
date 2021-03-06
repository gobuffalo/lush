// Code generated by github.com/gobuffalo/lush. DO NOT EDIT.
package goexamples

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/stretchr/testify/require"
)

func Test_mapExec(t *testing.T) {
	r := require.New(t)

	c := ast.NewContext(context.Background(), nil)

	s, err := lush.ParseFile("map.lush")
	r.NoError(err)
	r.True(Equal(c, s.Exec, mapExec))
}

var mapBResult *ast.Returned

func Benchmark_mapExec_Go(t *testing.B) {
	var r *ast.Returned

	for i := 0; i < t.N; i++ {
		c := ast.NewContext(context.Background(), nil)
		c.Imports.Store("fmt", builtins.NewFmt(ioutil.Discard))

		r, _ = mapExec(c)
	}
	mapBResult = r
}

func Benchmark_mapExec_Lush(t *testing.B) {
	var r *ast.Returned

	s, err := lush.ParseFile("map.lush")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < t.N; i++ {
		c := ast.NewContext(context.Background(), nil)
		c.Imports.Store("fmt", builtins.NewFmt(ioutil.Discard))

		r, _ = s.Exec(c)
	}
	mapBResult = r
}
