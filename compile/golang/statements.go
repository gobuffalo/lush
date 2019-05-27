package golang

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Compiler) astStatement(a ast.Statement) error {
	switch v := a.(type) {
	case ast.Script:
		return c.astScript(v)
	case ast.Return:
		return c.astReturn(v)
	case ast.Statements:
		return c.astStatements(v)
	case ast.Import:
		return c.astImport(v)
	case ast.Call:
		return c.astCall(v)
	case ast.Comment:
		_, err := fmt.Fprintf(c, "// %s\n", v.Value)
		return err
	case ast.If:
		return c.astIf(v)
	case *ast.Var:
		if err := c.astStatement(v.Name); err != nil {
			return err
		}
		fmt.Fprintf(c, " := ")
		if err := c.astStatement(v.Value); err != nil {
			return err
		}
		fmt.Fprintf(c, "\t_ = ")
		if err := c.astStatement(v.Name); err != nil {
			return err
		}
		fmt.Fprintf(c, "\n\n")
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

	default:
		fmt.Fprintln(c, a)
		return nil
		// panic(fmt.Sprintf("%T", a))
	}
	return nil
}

func (c Compiler) astStatements(a ast.Statements) error {
	for _, s := range a {
		if err := c.astStatement(s); err != nil {
			return err
		}
	}
	return nil
}
