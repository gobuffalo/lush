package match

import (
	"regexp"
	"strconv"
)

// Int ...
func Int(i int, pattern string) (bool, error) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return rx.MatchString(strconv.Itoa(i)), nil
}
