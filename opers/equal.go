package opers

import "github.com/google/go-cmp/cmp"

// Equalizer defines an interface to support the
// checking "equality" of one type to an another.
type Equalizer interface {
	Equal(b interface{}) (bool, error)
}

// Equal attempts to check "equality" of two types.
// Supports:
//	* Equalizer
//	* github.com/google/go-cmp/cmp
func Equal(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case Equalizer:
		return at.Equal(b)
	}

	res := cmp.Equal(a, b)

	return res, nil
}
