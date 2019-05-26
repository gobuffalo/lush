package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/modulus"
)

// Modulus attempts to take the "modulus" of type `a` with type `b`.
// Supports:
//	* int
//	* faces.Modulus
//	* faces.Int
func Modulus(a, b interface{}) (int, error) {
	switch at := a.(type) {
	case faces.Modulus:
		return at.Modulus(b)
	case int:
		return modulus.Int(at, b)
	case faces.Int:
		return modulus.Int(at.Int(), b)
	}

	return 0, modulus.Cant(a, b)
}
