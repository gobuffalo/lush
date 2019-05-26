package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/ne"
	"github.com/google/go-cmp/cmp"
)

// NotEqual `a == b`
// Supports:
//	* faces.NotEqual
//	* github.com/google/go-cmp/cmp
func NotEqual(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.NotEqual:
		return at.NotEqual(b)
	case map[string]interface{}:
		return ne.Map(at, b)
	case faces.Map:
		return ne.Map(at.Map(), b)
	case nil:
		return at != nil, nil
	}

	res := cmp.Equal(a, b)

	return !res, nil
}
