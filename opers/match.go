package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/match"
)

// Match will attempt to match the given regex pattern
// against the given type.
// Supports:
//	* string
//	* int
//	* float64
//	* bool
//	* fmt.Stringer
//	* faces.Match
//	* faces.Int
//	* faces.Float
//	* faces.Bool
func Match(i interface{}, pattern string) (bool, error) {
	switch s := i.(type) {
	case faces.Match:
		return s.Match(pattern)
	case string:
		return match.String(s, pattern)
	case fmt.Stringer:
		return match.String(s.String(), pattern)
	case int:
		return match.Int(s, pattern)
	case faces.Int:
		return match.Int(s.Int(), pattern)
	case float64:
		return match.Float(s, pattern)
	case faces.Float:
		return match.Float(s.Float(), pattern)
	case bool:
		return match.Bool(s, pattern)
	case faces.Bool:
		return match.Bool(s.Bool(), pattern)
	}
	return false, match.Cant(i, pattern)
}
