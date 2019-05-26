package match

import (
	"fmt"
	"regexp"
)

// Bool ...
func Bool(i bool, pattern string) (bool, error) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return rx.MatchString(fmt.Sprintf("%t", i)), nil
}
