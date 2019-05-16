package lush

import (
	"fmt"
	"io"

	"github.com/gobuffalo/lush/internal/parser"

	"github.com/gobuffalo/lush/ast"
)

// ParseFile parses the file identified by filename.
func ParseFile(filename string) (ast.Script, error) {
	return convert(parser.ParseFile(filename, parser.MetaOption))
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader) (ast.Script, error) {
	return convert(parser.ParseReader(filename, r, parser.MetaOption))

}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte) (ast.Script, error) {
	return convert(parser.Parse(filename, b, parser.MetaOption))
}

func convert(i interface{}, err error) (ast.Script, error) {
	s := ast.Script{}
	if err != nil {
		return s, err
	}
	sc, ok := i.(ast.Script)
	if !ok {
		return ast.Script{}, fmt.Errorf("expected ast.Script got %T", sc)
	}
	return sc, nil
}
