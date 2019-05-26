package match

import (
	"regexp"
	"strconv"
)

func Int(i int, pattern string) (bool, error) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return rx.MatchString(strconv.Itoa(i)), nil
}
