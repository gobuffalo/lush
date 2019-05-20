package quick

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func NewString(s string) ast.String {
	ns, _ := ast.NewString([]byte(fmt.Sprintf("%q", s)))
	return ns
}

func NewArray(i ...interface{}) ast.Array {
	arr, _ := ast.NewArray(i)
	return arr
}

func NewIdent(s string) ast.Ident {
	id, _ := ast.NewIdent([]byte(s))
	return id
}

func NewInteger(i int) ast.Integer {
	num, _ := ast.NewInteger(i)
	return num
}

func NewFloat(f float64) ast.Float {
	fl, _ := ast.NewFloat(f)
	return fl
}

func NewAssign(id ast.Ident, v ast.Statement) *ast.Assign {
	an, _ := ast.NewAssign(id, v)
	return an
}

func NewVar(id ast.Ident, v ast.Statement) *ast.Var {
	x, _ := ast.NewVar(id, v)
	return x
}

func NewAccess(id ast.Ident, i interface{}) ast.Access {
	acc, _ := ast.NewAccess(id, i)
	return acc
}

func NewBlock(s ...ast.Statement) *ast.Block {
	bl, _ := ast.NewBlock(s)
	return bl
}

func NewComment(s string) ast.Comment {
	return ast.Comment{Value: s}
}

func NewMap(m map[ast.Statement]interface{}) ast.Map {
	mp := ast.Map{
		Values: m,
	}

	return mp
}

func NewCall(n ast.Statement, y interface{}, args ast.Statements, b *ast.Block) ast.Call {
	c, _ := ast.NewCall(n, y, args, b)
	return c
}

func NewElse(b *ast.Block) ast.Else {
	el, _ := ast.NewElse(b)
	return el
}

func NewElseIf(fi ast.If) ast.ElseIf {
	el, _ := ast.NewElseIf(fi)
	return el
}

func NewIf(p ast.Statement, e ast.Expression, b *ast.Block, elsa ast.Statement) ast.If {
	i, _ := ast.NewIf(p, e, b, elsa)

	return i
}
