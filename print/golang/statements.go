package golang

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astStatement(a ast.Statement) error {
	// fmt.Printf("%T\n", a)
	switch v := a.(type) {
	case ast.Script:
		return c.astScript(v)
	case *ast.Let:
		return c.astLet(v)
	case ast.Return:
		return c.astReturn(v)
	case ast.Statements:
		return c.astStatements(v)
	case ast.Import:
		return c.astImport(v)
	case ast.Goroutine:
		return c.astGoroutine(v)
	case ast.Call:
		return c.astCall(v)
	case ast.Comment:
		return c.astComment(v)
	case ast.If:
		return c.astIf(v)
	case *ast.Var:
		return c.astVar(v)
	case ast.Range:
		return c.astRange(v)
	case ast.Ident:
		fmt.Fprint(c, v.Name)
		return nil
	case ast.String:
		fmt.Fprint(c, v)
		return nil
	case ast.Map:
		fmt.Fprintf(c, "%#v\n", v.Map())
		return nil
	case ast.Array:
		fmt.Fprintf(c, "%#v\n", v.Slice())
		return nil
	case ast.Func:
		return c.astFunc(v)
	default:
		fmt.Fprintln(c, a)
		return nil
	}
	return nil
}

func (c Printer) astStatements(a ast.Statements) error {
	for _, s := range a {
		if err := c.astStatement(s); err != nil {
			return err
		}
	}
	return nil
}
