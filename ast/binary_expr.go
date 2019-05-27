package ast

import (
	"errors"
	"fmt"
)

// BinaryExpr is an expression involving two operands. The operands can either
// be Unary expressions (e.g. !true) or another BinaryExpr for more complicated
// expressions (e.g. 4 + (6 / 2))
type BinaryExpr struct {
	Op  string
	LHS Visitable
	RHS Visitable
}

// Visit applies the operation of the BinaryExpr to the left and right subtrees
// (LHS and RHS) after recursively Visit-ing each subtree. This produces a
// depth-first evaluation order.
func (b BinaryExpr) Visit(c *Context) (interface{}, error) {
	lhsVal, err := b.LHS.Visit(c)
	if err != nil {
		return nil, err
	}

	rhsVal, err := b.RHS.Visit(c)
	if err != nil {
		return nil, err
	}

	switch b.Op {
	case "+":
		return lhsVal.(int) + rhsVal.(int), nil
	case "*":
		return lhsVal.(int) * rhsVal.(int), nil
	case "/":
		return lhsVal.(int) / rhsVal.(int), nil
	case "-":
		return lhsVal.(int) - rhsVal.(int), nil
	default:
		return nil, errors.New("Unsupported Operation " + b.Op)
	}
}

// String returns the infix formatted expression
func (b BinaryExpr) String() string {
	return fmt.Sprintf("(%#v %s %#v)", b.LHS, b.Op, b.RHS)
}
