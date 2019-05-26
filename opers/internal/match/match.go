package match

import "fmt"

// Cant ...
func Cant(a interface{}, pattern string) error {
	return fmt.Errorf("can not match %T with %q", a, pattern)
}
