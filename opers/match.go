package opers

import (
	"fmt"
	"regexp"

	"github.com/gobuffalo/lush/types"
)

// Matcher defines an interface to support the
// "matches" of one type from an another.
type Matcher interface {
	Match(string) (bool, error)
}

// Match will attempt to match the given regex pattern
// against the given type.
// Supports:
//	* string
//	* fmt.Stringer
//	* Matcher
func Match(i interface{}, pattern string) (bool, error) {
	fmt.Printf("### opers/match.go:23 pattern (%T) -> %q %+v\n", pattern, pattern, pattern)
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}

	switch s := i.(type) {
	case string:
		return rx.MatchString(s), nil
	case fmt.Stringer:
		return rx.MatchString(s.String()), nil
	case Matcher:
		return s.Match(pattern)
	default:
		// panic(fmt.Sprintf("%T", i))
		return rx.MatchString(types.Value(s)), nil
	}
	return false, nil
}
