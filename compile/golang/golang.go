package golang

import (
	"context"
	"io"
	"os"

	"github.com/gobuffalo/lush/ast"
)

type Compiler struct {
	context.Context
	io.Writer
}

var Default = Compiler{
	Context: context.Background(),
	Writer:  os.Stdout,
}

func (c Compiler) Compile(s ...ast.Statement) error {
	return c.astStatements(s)
}
