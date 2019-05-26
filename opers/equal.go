package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/equal"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Equal `a == b`
// Supports:
//	* faces.Equal
//	* github.com/google/go-cmp/cmp
func Equal(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.Equal:
		return at.Equal(b)
	case map[string]interface{}:
		return equal.Map(at, b)
	case faces.Map:
		return equal.Map(at.Map(), b)
	case nil:
		return at == nil, nil
	}

	res := cmp.Equal(a, b, cmpopts.IgnoreUnexported())

	return res, nil
}
