package goc

import (
	"github.com/gobuffalo/lush/ast"
)

func NewReturned(i ...interface{}) (*ast.Returned, error) {
	ret := ast.NewReturned(i)
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return &ret, nil
}
