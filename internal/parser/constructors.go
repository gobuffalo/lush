package parser

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/ast"
)

func newCurrent(c *current) (*ast.Current, error) {
	return &ast.Current{}, nil
}

func newCallExpr(c *current, head, tail, block interface{}) (ast.Node, error) {
	tailSlice, err := toII(tail)
	if err != nil {
		return nil, err
	}

	hn, err := toNode(head)
	if err != nil {
		return nil, err
	}

	if len(tailSlice) == 0 {
		return hn, nil
	}

	cur := ast.Call{
		Name: hn,
	}

	for _, t := range tailSlice {
		parts, err := toII(t)
		if err != nil {
			return nil, err
		}
		fname, fok := parts[1].(ast.Ident)
		if fok {
			cur.FName = fname
		}

		args, aok := parts[2].(ast.Nodes)
		if aok {
			cur.Arguments = args
		}

		if !fok && !aok {
			return nil, fmt.Errorf("method name and arguments missing")
		}
	}

	blk, err := toBlock(block)
	if err != nil {
		return nil, err
	}
	cur.Block = blk
	return cur, nil
}

func newArglist(c *current, head, tail interface{}) (ast.Nodes, error) {
	hn, err := toNode(head)
	if err != nil {
		return ast.Nodes([]ast.Node{}), nil
	}

	tailSlice, err := toII(tail)
	if err != nil {
		return nil, err
	}

	args := make([]ast.Node, 1, 1+len(tailSlice))
	args[0] = hn

	for _, t := range tailSlice {
		parts, err := toII(t)
		if err != nil {
			return nil, err
		}

		tn, err := toNode(parts[2])
		if err != nil {
			return nil, err
		}

		args = append(args, tn)
	}

	return ast.Nodes(args), nil
}

func newBinaryExpr(c *current, head, tail interface{}) (ast.Node, error) {
	hn, err := toNode(head)
	if err != nil {
		return nil, err
	}
	tailSlice, err := toII(tail)
	if err != nil {
		return nil, err
	}

	if len(tailSlice) == 0 {
		return hn, nil
	}

	cur := &ast.OpExpression{
		A: hn,
	}
	for _, parts := range tailSlice {
		tailParts, err := toII(parts)
		if err != nil {
			return nil, err
		}
		op, operand := tailParts[1], tailParts[3]

		cur.B = operand.(ast.Node)
		cur.Op = op.(string)
		cur = &ast.OpExpression{
			A: cur,
		}
	}
	return cur.A, nil
}

func newImport(c *current, s interface{}) (ast.Import, error) {
	n, ok := s.(ast.String)
	if !ok {
		return ast.Import{}, fmt.Errorf("expected ast.String got %T", s)
	}
	x, err := strconv.Unquote(n.Original)
	if err != nil {
		x = n.Original
	}
	i, err := ast.NewImport(x)
	i.Meta = meta(c)
	return i, err
}

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
	s, err := toNode(i)
	if err != nil {
		return ast.Call{}, err
	}

	args, err := toNodes(ax)
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

	idents, err := toIdentSlice(args)
	if err != nil {
		return ast.Range{}, err
	}
	ret, err = ast.NewRange(ni, idents, bl)
	if err != nil {
		return ret, err
	}

	ret.Meta = meta(c)
	return ret, nil
}

func newLHS(c *current, head interface{}, tail interface{}) ([]ast.Ident, error) {
	h, err := toIdent(head)
	if err != nil {
		return []ast.Ident{}, err
	}

	lhs := []ast.Ident{h}

	tailSlice, err := toII(tail)
	if err != nil {
		return []ast.Ident{}, err
	}

	for _, t := range tailSlice {
		tp, err := toII(t)
		if err != nil {
			return []ast.Ident{}, err
		}

		ident, err := toIdent(tp[3])
		if err != nil {
			return []ast.Ident{}, err
		}
		lhs = append(lhs, ident)
	}
	return lhs, nil
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

	idents, err := toIdentSlice(args)
	if err != nil {
		// no arguments, aka "infinite" for
		ret, err = ast.NewFor(ni, nil, bl)
		if err != nil {
			return ret, err
		}
		ret.Meta = meta(c)
		return ret, nil
	}

	ret, err = ast.NewFor(ni, idents, bl)
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
	states, err := toNodes(s)
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
	var ps ast.Node

	if p != nil {
		switch t := p.(type) {
		case []interface{}:
			st, err := toNodes(p)
			if err != nil {
				return ast.If{}, err
			}
			if len(st) > 0 {
				ps = st[0]
			}
		case ast.Node:
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

	cls, err := toNodes(elsa)
	if err != nil {
		return ast.If{}, err
	}

	var cl ast.Node
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

func newReturn(c *current, head, tail interface{}) (ret ast.Return, err error) {
	nodes, err := toNodesFromList(head, tail)
	if err != nil {
		return ast.Return{}, err
	}
	ret, err = ast.NewReturn(nodes)
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

	sv, err := toNode(v)
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
	in, err := toNode(n)
	if err != nil {
		return nil, err
	}

	sv, err := toNode(v)
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

	sv, err := toNode(v)
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
	sa, ok := a.(ast.Node)
	if !ok {
		return nil, fmt.Errorf("expected ast.Node got %T", a)
	}
	sb, ok := b.(ast.Node)
	if !ok {
		return nil, fmt.Errorf("expected ast.Node got %T", b)
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

func newNoop(c *current) (ast.Noop, error) {
	n, err := ast.NewNoop(c.text)
	n.Meta = meta(c)
	return n, err
}
