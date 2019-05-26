package sub

import "fmt"

// Cant ...
func Cant(at, b interface{}) error {
	return fmt.Errorf("can not subtract %T and %T", at, b)
}
