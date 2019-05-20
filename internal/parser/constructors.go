package parser

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/ast"
)

func newString(c *current) (ast.String, error) {
	s, err := ast.NewString(c.text)
	if err != nil {
		return s, err
	}
	s.Meta = meta(c)

	return s, nil
}

func newAccess(c *current, i interface{}, k interface{}) (ret ast.Access, err error) {
	id, err := toIdent(i)
	if err != nil {
		return ast.Access{}, err
	}

	ret, err = ast.NewAccess(id, k)
	if err != nil {
		return ret, err
	}
	ret.Meta = meta(c)
	return ret, nil
}

func newComment(c *current, b []byte) (ret ast.Comment, err error) {
	ret, err = ast.NewComment(b)
	if err != nil {
		return ret, err
	}
	ret.Meta = meta(c)
	return ret, nil
}

func newMap(c *current, vals interface{}) (ret ast.Map, err error) {
	ret, err = ast.NewMap(vals)
	if err != nil {
		return ret, err
	}
	ret.Meta = meta(c)
	return ret, nil
}

func newBool(c *current, b []byte) (ret ast.Bool, err error) {
	ret, err = ast.NewBool(b)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newCall(c *current, i interface{}, y interface{}, ax interface{}, b interface{}) (ret ast.Call, err error) {
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

	ret, err = ast.NewCall(s, y, args, bl)
	if err != nil {
		return ret, err
	}
	ret.Meta = meta(c)
	return ret, nil
}

func newFunc(c *current, ax interface{}, b interface{}) (ret ast.Func, err error) {
	bl, err := toBlock(b)
	if err != nil {
		return ast.Func{}, err
	}
	ret, err = ast.NewFunc(ax, bl)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newRange(c *current, n interface{}, args interface{}, b interface{}) (ret ast.Range, err error) {
	bl, err := toBlock(b)
	if err != nil {
		return ast.Range{}, err
	}
	ni, err := toExecString(n)
	if err != nil {
		return ast.Range{}, err
	}

	ret, err = ast.NewRange(ni, args, bl)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newFor(c *current, n interface{}, args interface{}, b interface{}) (ret ast.For, err error) {
	bl, err := toBlock(b)
	if err != nil {
		return ast.For{}, err
	}
	ni, err := toExecString(n)
	if err != nil {
		return ast.For{}, err
	}
	ret, err = ast.NewFor(ni, args, bl)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newIdent(c *current, b []byte) (ret ast.Ident, err error) {
	ret, err = ast.NewIdent(b)
	if err != nil {
		return ret, err
	}
	ret.Meta = meta(c)
	return ret, nil
}

func newBlock(c *current, s interface{}) (ret *ast.Block, err error) {
	states, err := toStatements(s)
	if err != nil {
		return nil, err
	}
	ret, err = ast.NewBlock(states)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newIf(c *current, p interface{}, e interface{}, b interface{}, elsa interface{}) (ret ast.If, err error) {
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

	ret, err = ast.NewIf(ps, es, bl, cl)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newElseIf(c *current, i interface{}) (ret ast.ElseIf, err error) {
	fi, err := toIf(i)
	if err != nil {
		return ast.ElseIf{}, err
	}
	ret, err = ast.NewElseIf(fi)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newReturn(c *current, i interface{}) (ret ast.Return, err error) {
	s, err := toStatements(i)
	if err != nil {
		return ast.Return{}, err
	}
	ret, err = ast.NewReturn(s)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newFloat(c *current, b []byte) (ret ast.Float, err error) {
	f, err := strconv.ParseFloat(string(b), 64)
	fl, err := ast.NewFloat(f)
	fl.Meta = meta(c)

	return fl, err
}

func newInteger(c *current, b []byte) (ret ast.Integer, err error) {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return ast.Integer{}, err
	}

	in, err := ast.NewInteger(i)
	if err != nil {
		return ast.Integer{}, err
	}
	in.Meta = meta(c)
	return in, nil
}

func newLet(c *current, n, v interface{}) (ret *ast.Let, err error) {
	in, ok := n.(ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ast.Ident got %T", n)
	}

	sv, err := toStatement(v)
	if err != nil {
		return nil, err
	}
	ret, err = ast.NewLet(in, sv)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newAssign(c *current, n, v interface{}) (ret *ast.Assign, err error) {
	in, ok := n.(ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ast.Ident got %T", n)
	}

	sv, err := toStatement(v)
	if err != nil {
		return nil, err
	}
	ret, err = ast.NewAssign(in, sv)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newVar(c *current, n, v interface{}) (ret *ast.Var, err error) {
	in, ok := n.(ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected ast.Ident got %T", n)
	}

	sv, err := toStatement(v)
	if err != nil {
		return nil, err
	}
	ret, err = ast.NewVar(in, sv)
	if err != nil {
		return nil, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newOpExpression(c *current, a, op, b interface{}) (ret *ast.OpExpression, err error) {
	// defer setMeta(&ret, c)
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
	ret, err = ast.NewOpExpression(sa, sop, sb)
	if err != nil {
		return nil, err
	}
	ret.Meta = meta(c)
	return ret, nil
}

func newPopExpression(c *current, a, op, b interface{}) (ret *ast.OpExpression, err error) {
	ope, err := newOpExpression(c, a, op, b)
	if err != nil {
		return nil, err
	}

	ret, err = ast.NewPopExpression(ope.A, ope.Op, ope.B)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newElse(c *current, i interface{}) (ret ast.Else, err error) {
	b, err := toBlock(i)
	if err != nil {
		return ast.Else{}, err
	}
	ret, err = ast.NewElse(b)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newArray(c *current, i interface{}) (ret ast.Array, err error) {
	ii, err := toII(i)
	if err != nil {
		return ast.Array{}, err
	}

	ret, err = ast.NewArray(ii)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}
