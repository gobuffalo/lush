package match

import (
	"regexp"
)

func String(s string, pattern string) (bool, error) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return rx.MatchString(s), nil
}
