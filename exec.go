package lush

import (
	"io"

	"github.com/gobuffalo/lush/ast"
)

// Exec a script using the specified Runtime.
func Exec(c *ast.Runtime, s ast.Script) (*ast.Returned, error) {
	return s.Exec(c)
}

// ExecFile will parse and then execute the specified file.
func ExecFile(c *ast.Runtime, filename string) (*ast.Returned, error) {
	s, err := ParseFile(filename)
	if err != nil {
		return nil, err
	}
	return Exec(c, s)
}

// ExecReader will parse the given reader and then execute it.
func ExecReader(c *ast.Runtime, filename string, r io.Reader) (*ast.Returned, error) {
	s, err := ParseReader(filename, r)
	if err != nil {
		return nil, err
	}
	return Exec(c, s)
}
