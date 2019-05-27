package golang

import (
	"context"
	"io"
	"os"

	"github.com/gobuffalo/lush/ast"
)

type Printer struct {
	context.Context
	io.Writer
}

var Default = Printer{
	Context: context.Background(),
	Writer:  os.Stdout,
}

func (c Printer) Print(s ...ast.Node) error {
	return c.astNodes(s)
}
