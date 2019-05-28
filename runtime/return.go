package runtime

import (
	"github.com/gobuffalo/lush"
	"github.com/gobuffalo/lush/ast"
)

var Current = Runtime{
	Version: lush.Version,
}

type Runtime struct {
	Version string
}

func (Runtime) NewReturned(i ...interface{}) (*ast.Returned, error) {
	ret := ast.NewReturned(i)
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return &ret, nil
}
