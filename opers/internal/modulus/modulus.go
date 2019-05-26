package modulus

import "fmt"

// Cant ...
func Cant(at, b interface{}) error {
	return fmt.Errorf("can not take the modulus of %T and %T", at, b)
}
