package add

import "fmt"

// Cant ...
func Cant(at, b interface{}) error {
	return fmt.Errorf("can not add %T (%v) and %T (%v)", at, at, b, b)
}
