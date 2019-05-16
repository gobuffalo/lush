package parser

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/ast"
)

func newAccess(c *current, i interface{}, k interface{}) (ret ast.Access, err error) {
	defer setMeta(&ret, c)
	id, err := toIdent(i)
	if err != nil {
		return ast.Access{}, err
	}
	return ast.NewAccess(id, k)
}

func newComment(c *current, b []byte) (ret ast.Comment, err error) {
	defer setMeta(&ret, c)
	return ast.NewComment(b)
}

func newMap(c *current, vals interface{}) (ret ast.Map, err error) {
	defer setMeta(&ret, c)
	return ast.NewMap(vals)
}

func newBool(c *current, b []byte) (ret ast.Bool, err error) {
	defer setMeta(&ret, c)
	return ast.NewBool(b)
}

func newCall(c *current, i interface{}, y interface{}, ax interface{}, b interface{}) (ret ast.Call, err error) {
	defer setMeta(&ret, c)
	s, err := toStatement(i)
	if err != nil {
		return ast.Call{}, err
	}

	args, err := toStatements(ax)
	if err != nil {
		return ast.Call{}, err
	}

	bl, err := toBlock(b)
	if err != nil {
		return ast.Call{}, err
	}

	return ast.NewCall(s, y, args, bl)
}

func newFunc(c *current, ax interface{}, b interface{}) (ret ast.Func, err error) {
	defer setMeta(&ret, c)
	bl, err := toBlock(b)
	if err != nil {
		return ast.Func{}, err
	}
	return ast.NewFunc(ax, bl)
}

func newRange(c *current, n interface{}, args interface{}, b interface{}) (ret ast.Range, err error) {
	defer setMeta(&ret, c)
	bl, err := toBlock(b)
	if err != nil {
		return ast.Range{}, err
	}
	ni, err := toExecString(n)
	if err != nil {
		return ast.Range{}, err
	}
	return ast.NewRange(ni, args, bl)
}

func newFor(c *current, n interface{}, args interface{}, b interface{}) (ret ast.For, err error) {
	defer setMeta(&ret, c)
	bl, err := toBlock(b)
	if err != nil {
		return ast.For{}, err
	}
	ni, err := toExecString(n)
	if err != nil {
		return ast.For{}, err
	}
	return ast.NewFor(ni, args, bl)
}

func newIdent(c *current, b []byte) (ret ast.Ident, err error) {
	defer setMeta(&ret, c)
	return ast.NewIdent(b)
}

func newBlock(c *current, s interface{}) (ret *ast.Block, err error) {
	defer setMeta(&ret, c)
	states, err := toStatements(s)
	if err != nil {
		return nil, err
	}
	return ast.NewBlock(states)
}

func newIf(c *current, p interface{}, e interface{}, b interface{}, elsa interface{}) (ret ast.If, err error) {
	defer setMeta(&ret, c)
	var ps ast.Statement

	if p != nil {
		switch t := p.(type) {
		case []interface{}:
			st, err := toStatements(p)
			if err != nil {
				return ast.If{}, err
			}
			if len(st) > 0 {
				ps = st[0]
			}
		case ast.Statement:
			ps = t
		}
	}

	bl, err := toBlock(b)
	if err != nil {
		return ast.If{}, err
	}

	es, err := toExpression(e)
	if err != nil {
		return ast.If{}, err
	}

	cls, err := toStatements(elsa)
	if err != nil {
		return ast.If{}, err
	}

	var cl ast.Statement
	if len(cls) > 0 {
		cl = cls[0]
	}

	return ast.NewIf(ps, es, bl, cl)
}

func newElseIf(c *current, i interface{}) (ret ast.ElseIf, err error) {
	defer setMeta(&ret, c)
	fi, err := toIf(i)
	if err != nil {
		return ast.ElseIf{}, err
	}
	return ast.NewElseIf(fi)
}

func newReturn(c *current, i interface{}) (ret ast.Return, err error) {
	defer setMeta(&ret, c)
	s, err := toStatements(i)
	if err != nil {
		return ast.Return{}, err
	}
	return ast.NewReturn(s)
}

func newFloat(c *current, b []byte) (ret ast.Float, err error) {
	defer setMeta(&ret, c)
	f, err := strconv.ParseFloat(string(b), 64)
	return ast.Float(f), err
}

func newInteger(c *current, b []byte) (ret ast.Integer, err error) {
	defer setMeta(&ret, c)
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return ast.Integer(i), err
	}
	return ast.Integer(i), nil
}

func newLet(c *current, n, v interface{}) (ret *ast.Let, err error) {
	defer setMeta(&ret, c)
	in, ok := n.(ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ast.Ident got %T", n)
	}

	sv, err := toStatement(v)
	if err != nil {
		return nil, err
	}
	return ast.NewLet(in, sv)
}

func newAssign(c *current, n, v interface{}) (ret *ast.Assign, err error) {
	defer setMeta(&ret, c)
	in, ok := n.(ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ast.Ident got %T", n)
	}

	sv, err := toStatement(v)
	if err != nil {
		return nil, err
	}
	return ast.NewAssign(in, sv)
}

func newVar(c *current, n, v interface{}) (ret *ast.Var, err error) {
	defer setMeta(&ret, c)
	in, ok := n.(ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ast.Ident got %T", n)
	}

	sv, err := toStatement(v)
	if err != nil {
		return nil, err
	}
	return ast.NewVar(in, sv)
}

func newOpExpression(c *current, a, op, b interface{}) (ret *ast.OpExpression, err error) {
	defer setMeta(&ret, c)
	sa, ok := a.(ast.Statement)
	if !ok {
		return nil, fmt.Errorf("expected ast.Statement got %T", a)
	}
	sb, ok := b.(ast.Statement)
	if !ok {
		return nil, fmt.Errorf("expected ast.Statement got %T", b)
	}
	sop, ok := op.(string)
	if !ok {
		return nil, fmt.Errorf("expected string got %T", op)
	}
	return ast.NewOpExpression(sa, sop, sb)
}

func newPopExpression(c *current, a, op, b interface{}) (ret *ast.OpExpression, err error) {
	defer setMeta(&ret, c)
	ope, err := newOpExpression(c, a, op, b)
	if err != nil {
		return nil, err
	}
	return ast.NewPopExpression(ope.A, ope.Op, ope.B)
}

func newElse(c *current, i interface{}) (ret ast.Else, err error) {
	defer setMeta(&ret, c)
	b, err := toBlock(i)
	if err != nil {
		return ast.Else{}, err
	}
	return ast.NewElse(b)
}

func newArray(c *current, i interface{}) (ret ast.Array, err error) {
	defer setMeta(&ret, c)
	ii, err := toII(i)
	if err != nil {
		return ast.Array{}, err
	}
	return ast.NewArray(ii)
}
