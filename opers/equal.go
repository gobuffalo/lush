package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/google/go-cmp/cmp"
)

// Equal attempts to check "equality" of two types.
// Supports:
//	* faces.Equal
//	* github.com/google/go-cmp/cmp
func Equal(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.Equal:
		return at.Equal(b)
	}

	res := cmp.Equal(a, b)

	return res, nil
}
