package parser

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func toIdent(i interface{}) (ast.Ident, error) {
	id, ok := i.(ast.Ident)
	if !ok {
		return ast.Ident{}, fmt.Errorf("expected ast.Ident got %T", i)
	}
	return id, nil
}

func toIdentSlice(i interface{}) ([]ast.Ident, error) {
	ii, ok := i.([]ast.Ident)
	if !ok {
		return nil, fmt.Errorf("expected []ast.Ident, got %T", i)
	}
	return ii, nil
}

func toExecString(i interface{}) (ast.ExecableNode, error) {
	if i == nil {
		return nil, nil
	}
	n, err := toNode(i)
	if err != nil {
		return nil, err
	}

	in, ok := n.(ast.ExecableNode)
	if !ok {
		return nil, fmt.Errorf("expected ExecStringer, got %T", i)
	}
	return in, nil
}

func toExpression(e interface{}) (ast.Expression, error) {
	if e == nil {
		return nil, nil
	}
	if ii, ok := e.([]interface{}); ok {
		for _, s := range ii {
			if _, ok := s.(ast.Noop); ok {
				continue
			}
			e = s
			break
		}
	}

	es, ok := e.(ast.Expression)
	if !ok {
		return nil, fmt.Errorf("expected ast.Expression got %T", e)
	}
	return es, nil
}

func toBlock(i interface{}) (*ast.Block, error) {
	if i == nil {
		return nil, nil
	}
	eb, ok := i.(*ast.Block)
	if !ok {
		return nil, fmt.Errorf("expected *ast.Block got %T", i)
	}
	return eb, nil
}

func toII(i interface{}) ([]interface{}, error) {
	ii, ok := i.([]interface{})
	if !ok {
		return ii, fmt.Errorf("expected []interface{} got %T", i)
	}
	return ii, nil
}

func toNode(i interface{}) (ast.Node, error) {
	ii, ok := i.(ast.Node)
	if !ok {
		return ii, fmt.Errorf("expected ast.Node got %T", i)
	}
	return ii, nil
}

func toNodes(i interface{}) (ast.Nodes, error) {
	nds, ok := i.(ast.Nodes)
	if ok {
		return nds, nil
	}

	ii, err := toII(i)
	if err != nil {
		return nil, err
	}
	var states ast.Nodes

	for _, s := range ii {
		st, err := toNode(s)
		if err != nil {
			return nil, err
		}
		states = append(states, st)
	}
	return states, nil
}

func toNodesFromList(head, tail interface{}) (ast.Nodes, error) {
	hn, err := toNode(head)
	if err != nil {
		return nil, err
	}
	tailSlice, err := toII(tail)
	if err != nil {
		return nil, err
	}
	out := make([]ast.Node, 1, len(tailSlice)+1)
	out[0] = hn

	for _, p := range tailSlice {
		parts, err := toII(p)
		if err != nil {
			return ast.Nodes([]ast.Node{}), err
		}
		node, err := toNode(parts[2])
		if err != nil {
			return ast.Nodes([]ast.Node{}), err
		}
		out = append(out, node)
	}

	return ast.Nodes(out), nil
}

func toIf(i interface{}) (ast.If, error) {
	fi, ok := i.(ast.If)
	if !ok {
		return ast.If{}, fmt.Errorf("expected ast.If got %T", i)
	}
	return fi, nil
}
