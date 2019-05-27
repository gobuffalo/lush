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

func NewAssign(id ast.Ident, v ast.Node) *ast.Assign {
	an, _ := ast.NewAssign(id, v)
	return an
}

func NewVar(id ast.VarRef, v ast.Node) *ast.Var {
	x, _ := ast.NewVar(id, v)
	return x
}

func NewVarRef(name string) ast.VarRef {
	return ast.VarRef{
		Name: name,
	}
}

func NewAccess(id ast.Ident, i interface{}) ast.Access {
	acc, _ := ast.NewAccess(id, i)
	return acc
}

func NewBlock(s ...ast.Node) *ast.Block {
	bl, _ := ast.NewBlock(s...)
	return bl
}

func NewComment(s string) ast.Comment {
	return ast.Comment{Value: s}
}

func NewMap(m map[string]interface{}) ast.Map {
	mp := ast.Map{
		Values: m,
	}

	return mp
}

func NewCall(n ast.Node, y interface{}, args ast.Nodes, b *ast.Block) ast.Call {
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

func NewIf(p ast.Node, e ast.Expression, b *ast.Block, elsa ast.Node) ast.If {
	i, _ := ast.NewIf(p, e, b, elsa)

	return i
}

func NewFor(n ast.VisitableNode, args interface{}, b *ast.Block) ast.For {
	f, _ := ast.NewFor(n, args, b)

	return f
}

func NewRange(n ast.VisitableNode, args interface{}, b *ast.Block) ast.Range {
	f, _ := ast.NewRange(n, args, b)

	return f
}

func NewFunc(ax interface{}, b *ast.Block) ast.Func {
	f, _ := ast.NewFunc(ax, b)
	return f
}

func NewLet(name ast.Ident, value ast.Node) *ast.Let {
	l, _ := ast.NewLet(name, value)

	return l
}

func NewOpExpression(a ast.Node, op string, b ast.Node) *ast.OpExpression {
	o, _ := ast.NewOpExpression(a, op, b)

	return o
}

func NewReturn(s ...ast.Node) ast.Return {
	r, _ := ast.NewReturn(s)
	return r
}

func NewImport(s string) ast.Import {
	i, _ := ast.NewImport(s)
	return i
}
