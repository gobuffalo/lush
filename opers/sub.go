package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/sub"
)

// Sub `a - b`
// Supports:
//	* int
//	* float64
//	* map[string]interface{}
//	* faces.Sub
//	* faces.Float
//	* faces.Int
//	* faces.Map
func Sub(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case faces.Sub:
		return at.Sub(b)
	case int:
		return sub.Int(at, b)
	case float64:
		return sub.Float(at, b)
	case faces.Int:
		return sub.Int(at.Int(), b)
	case faces.Float:
		return sub.Float(at.Float(), b)
	case map[string]interface{}:
		return sub.Map(at, b)
	case faces.Map:
		return sub.Map(at.Map(), b)
	}

	return nil, sub.Cant(a, b)
}
