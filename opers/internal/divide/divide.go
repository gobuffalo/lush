package divide

import "fmt"

// Cant ...
func Cant(at, b interface{}) error {
	return fmt.Errorf("can not divide %T and %T", at, b)
}
