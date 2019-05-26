package match

import (
	"fmt"
	"regexp"
)

func Float(i float64, pattern string) (bool, error) {
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return rx.MatchString(fmt.Sprintf("%f", i)), nil
}
