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

func toExecString(i interface{}) (ast.VisitableStatement, error) {
	if i == nil {
		return nil, nil
	}
	n, err := toStatement(i)
	if err != nil {
		return nil, err
	}

	in, ok := n.(ast.VisitableStatement)
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

func toStatement(i interface{}) (ast.Statement, error) {
	ii, ok := i.(ast.Statement)
	if !ok {
		return ii, fmt.Errorf("expected ast.Statement got %T", i)
	}
	return ii, nil
}

func toStatements(i interface{}) (ast.Statements, error) {
	ii, err := toII(i)
	if err != nil {
		return nil, err
	}
	var states ast.Statements

	for _, s := range ii {
		st, err := toStatement(s)
		if err != nil {
			return nil, err
		}
		states = append(states, st)
	}
	return states, nil
}

func toIf(i interface{}) (ast.If, error) {
	fi, ok := i.(ast.If)
	if !ok {
		return ast.If{}, fmt.Errorf("expected ast.If got %T", i)
	}
	return fi, nil
}
