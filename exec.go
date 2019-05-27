package lush

import (
	"io"

	"github.com/gobuffalo/lush/ast"
)

// Visit a script using the specified Context.
func Exec(c *ast.Context, s ast.Script) (*ast.Returned, error) {
	return s.Exec(c)
}

// ExecFile will parse and then execute the specified file.
func ExecFile(c *ast.Context, filename string) (*ast.Returned, error) {
	s, err := ParseFile(filename)
	if err != nil {
		return nil, err
	}
	return Exec(c, s)
}

// ExecReader will parse the given reader and then execute it.
func ExecReader(c *ast.Context, filename string, r io.Reader) (*ast.Returned, error) {
	s, err := ParseReader(filename, r)
	if err != nil {
		return nil, err
	}
	return Exec(c, s)
}
