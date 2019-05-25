package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/lessthan"
)

func LessThan(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.LessThan:
		return at.LessThan(b)
	case int:
		return lessthan.Int(at, b)
	case float64:
		return lessthan.Float(at, b)
	case string:
		return lessthan.String(at, b)
	case fmt.Stringer:
		return lessthan.String(at.String(), b)
	case faces.Int:
		return lessthan.Int(at.Int(), b)
	case faces.Float:
		return lessthan.Float(at.Float(), b)
	}
	return false, fmt.Errorf("can not compare %T with %T", a, b)
}
