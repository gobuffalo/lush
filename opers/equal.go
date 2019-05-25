package opers

import "github.com/google/go-cmp/cmp"

type Equalizer interface {
	Equal(b interface{}) (bool, error)
}

func Equal(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case Equalizer:
		return at.Equal(b)
	}

	res := cmp.Equal(a, b)

	return res, nil
}
