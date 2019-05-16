package ast_test

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/internal/parser"
)

var (
	abc     = []interface{}{"a", "b", "c"}
	beatles = map[string]string{
		"john":   "guitar",
		"paul":   "bass",
		"george": "guitar",
		"ringo":  "drums",
	}
)

func NewContext() *ast.Context {
	var w io.Writer = ioutil.Discard
	if testing.Verbose() {
		w = os.Stdout
	}
	return ast.NewContext(context.Background(), w)
}

func exec(in string, c *ast.Context) (*ast.Returned, error) {
	st, err := parse(in)
	if err != nil {
		return nil, err
	}
	return st.Exec(c)
}

func parse(in string) (ast.Script, error) {
	p, err := parser.Parse("x.plush", []byte(in))
	if err != nil {
		return ast.Script{}, err
	}
	n, ok := p.(ast.Script)
	if !ok {
		return n, errors.New("not a node!")
	}
	return n, nil
}

func newString(in string) ast.String {
	s, err := ast.NewString([]byte(in))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
