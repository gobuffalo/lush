package lessthan

import "fmt"

// Cant ...
func Cant(at, b interface{}) error {
	return fmt.Errorf("can not compare %T and %T", at, b)
}
